package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	s "trustgrid.io/swaggerbot/pkg/swagger"
)

var validationFailure *s.Schema

var params struct {
	domainName  *s.Param
	nodeID      *s.Param
	clusterFQDN *s.Param
}

var examples struct {
	nodeID      string
	appID       string
	sessionID   string
	nodeName    string
	fqdn        string
	clusterFQDN string
}

func addAlertAPI(api *s.API) {
	p := api.Path("/alert").Permission("alerts::read").Produces("application/json").Tag("Alerts")

	alert := api.Model("Alert").
		Prop("uid", s.NewSchema("Unique ID of alert", s.S_String, s.S_Example("2DqxLdknjWxEkGt474d2Cstsa1O"))).
		Prop("type", s.NewSchema("Type of alert", s.S_String, s.S_Example("Data Plane Disruption"))).
		Prop("node", s.NewSchema("Node name that generated the alert", s.S_String, s.S_Example(examples.nodeName))).
		Prop("node_id", s.NewSchema("Node ID that generated the alert", s.S_String, s.S_Example(examples.nodeID))).
		Prop("timestamp", s.NewSchema("Unix timestamp of alert", s.S_Number, s.S_Example(1661440360))).
		Prop("alert_message", s.NewSchema("More information about the alert", s.S_String, s.S_Example("Node mynode abnormally disconnected"))).
		Ref()

	p.Get("List events, newest first").
		Param(s.NewParam("timestamp", "Cursor returned from previous request", s.P_Number)).
		Param(s.NewParam("limit", "Limit number of alerts to return", s.P_Number)).
		Response(200, "OK", s.NewArraySchema(alert))

	p = p.PathParam(params.nodeID)
	p.Get("List events for a node, newest first").Response(200, "OK", s.NewArraySchema(alert))
}

func addV2AlertAPI(api *s.API) {
	p := api.Path("/v2/alert").Permission("alerts::read").Produces("application/json").Tag("Alerts")

	alert := api.Model("Alert V2").
		Prop("uid", s.NewSchema("Unique ID of alert", s.S_String, s.S_Example("2DqxLdknjWxEkGt474d2Cstsa1O"))).
		Prop("eventType", s.NewSchema("Type of alert", s.S_String, s.S_Example("Data Plane Disruption"))).
		Prop("nodeName", s.NewSchema("Node name that generated the alert", s.S_String, s.S_Example(examples.nodeName))).
		Prop("nodeId", s.NewSchema("Node ID that generated the alert", s.S_String, s.S_Example(examples.nodeID))).
		Prop("timestamp", s.NewSchema("Unix timestamp of alert", s.S_Number, s.S_Example(1661440360))).
		Prop("message", s.NewSchema("More information about the alert", s.S_String, s.S_Example("Node mynode abnormally disconnected"))).
		Ref()

	p.Get("List alerts, newest first").
		Response(200, "OK", s.NewArraySchema(alert))

	p = p.PathParam(params.nodeID)
	p.Get("List alerts for a node, newest first").Response(200, "OK", s.NewArraySchema(alert))

	p = p.PathParam(s.NewParam("alertType", "Alert type, eg Node Disconnect", s.P_Path, s.P_Required))
	p.Delete("Resolve alert").Response(200, "OK", nil)
}

func addTagsAPI(api *s.API) {
	p := api.Path("/tags").Produces("application/json").Tag("Tags")

	tag := api.Model("Tag").
		Prop("name", s.NewSchema("Tag name", s.S_String, s.S_Example("prod_status"))).
		Prop("values", s.NewArraySchema(s.NewSchema("Tag values", s.S_String))).
		Prop("created_at", s.NewSchema("Unix timestamp when tag was created", s.S_Number)).
		Ref()

	p.Get("List tags").Response(200, "OK", s.NewArraySchema(tag))

	p = p.PathParam(s.NewParam("tagName", "Tag name", s.P_Path, s.P_Required))
	p.Get("Get a tag by name").
		Response(200, "OK", tag).
		Response(404, "Tag not found", nil)
}

func addCertificatesAPI(api *s.API) {
	p := api.Path("/v2/certificates").Produces("application/json").Consumes("application/json").Tag("Certificates")

	cert := api.Model("Certificate").
		Prop("fqdn", s.NewSchema("FQDN of certificate", s.S_String, s.S_Example(examples.fqdn))).
		Prop("expiresAt", s.NewSchema("Unix timestamp when certificate expires", s.S_Number)).
		Prop("warning", s.NewSchema("Warning message", s.S_String)).
		Ref()

	p.Get("List certificates").Permission("certificates::read").Response(200, "OK", s.NewArraySchema(cert))

	certParts := api.Model("CertificateParts").
		Prop("private_key", s.NewSchema("Private key for the certificate in PEM format", s.S_String, s.S_Required)).
		Prop("chain", s.NewSchema("Certificate chain in PEM format", s.S_String, s.S_Required)).
		Ref()

	certParam := s.NewParam("CertParts", "Certificate parts", s.P_Body, s.P_Schema(certParts), s.P_Required)

	p.Post("Create a certificate").
		Permission("certificates::modify").
		Param(certParam).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	p = p.PathParam(s.NewParam("fqdn", "FQDN of certificate", s.P_Path, s.P_Required))
	p.Permission("certificates::modify")
	p.Put("Update a certificate").
		Param(certParam).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	p.Delete("Delete a certificate").Response(200, "OK", nil)
}

func addDomainAPI(api *s.API) {
	p := api.Path("/domain").Produces("application/json").Consumes("application/json").Tag("Domain")
	p = p.PathParam(params.domainName)

	domain := s.NewSchema("Domain")
	domain.Prop("config", s.NewSchema("Domain configuration"))

	p.Get("Get domain details").Permission("domains::read").Response(200, "OK", domain).Response(404, "Not found", nil)

	p.Path("/config/alert").
		Put("Update alert configuration").
		Permission("domains::configure:thresholds").
		Param(s.NewParam("alert", "Alert configuration", s.P_Body, s.P_Schema(s.NewSchema("Alert configuration")), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	p.Path("/config/apigw").
		Put("Update API GW configuration").
		Permission("domains::configure:gateway").
		Param(s.NewParam("apigw", "APIGW configuration", s.P_Body, s.P_Schema(s.NewSchema("APIGW configuration")), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
}

func addApplicationAPI(api *s.API) {
	p := api.Path("/v2/application").Produces("application/json").Consumes("application/json").Tag("Applications")

	a := api.Model("Application").
		Prop("uid", s.NewSchema("Unique ID of application", s.S_String)).
		Prop("appType", s.NewSchema("Type of application", s.S_String, s.S_Enum("web", "remote", "wireguard"))).
		Prop("hostname", s.NewSchema("Hostname of application", s.S_String)).
		Prop("name", s.NewSchema("Name of application", s.S_String)).
		Prop("gatewayNode", s.NewSchema("Gateway node for the application", s.S_String)).
		Prop("edgeNode", s.NewSchema("Edge node for the application", s.S_String)).
		Prop("description", s.NewSchema("", s.S_String)).
		Ref()

	p.Get("List applications your user can access").Response(200, "OK", s.NewArraySchema(a))

	p.Post("Create an application").
		Permission("applications::modify").
		Param(s.NewParam("application", "Application configuration", s.P_Body, s.P_Schema(a), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	p.Path("/all").
		Permission("applications::read").
		Get("List applications").
		Response(200, "OK", s.NewArraySchema(a))

	id := p.PathParam(s.NewParam("applicationID", "Unique ID of application", s.P_Path, s.P_Required))
	id.Get("Get an application").
		Permission("applications::read").
		Response(200, "OK", a).
		Response(404, "Not found", nil)
	id.Delete("Delete an application").
		Permission("applications::modify").
		Response(200, "OK", nil).
		Response(404, "Not found", nil)
	id.Put("Update an application").
		Permission("applications::modify").
		Param(s.NewParam("application", "Application configuration", s.P_Body, s.P_Schema(a), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	criteria := api.Model("Criteria").
		Prop("emails", s.NewArraySchema(s.NewSchema("Email addresses", s.S_String))).
		Prop("emailsEndingIn", s.NewArraySchema(s.NewSchema("Email suffixes", s.S_String, s.S_Example("@trustgrid.io")))).
		Prop("everyone", s.NewSchema("Everyone", s.S_Boolean)).
		Prop("country", s.NewArraySchema(s.NewSchema("Country codes", s.S_String, s.S_Example("US")))).
		Prop("idpGroups", s.NewArraySchema(s.NewSchema("IDP Group IDs", s.S_String))).
		Prop("ipRanges", s.NewArraySchema(s.NewSchema("IP Ranges", s.S_String))).
		Ref()

	axsRule := api.Model("Application Access Rule").
		Prop("action", s.NewSchema("Action", s.S_String, s.S_Enum("allow", "block"), s.S_Required)).
		Prop("exceptions", criteria).
		Prop("includes", criteria).
		Prop("requires", criteria).
		Prop("name", s.NewSchema("Name", s.S_String, s.S_Required)).
		Ref()

	axsPolicy := id.Path("/access-policy")

	axsPolicy.Get("List access rules").
		Permission("applications::read").
		Response(200, "OK", s.NewArraySchema(axsRule))
	axsPolicy.Post("Create an access rule").
		Permission("applications::modify").
		Param(s.NewParam("access rule", "Access rule", s.P_Body, s.P_Schema(axsRule), s.P_Required)).
		Response(200, "OK", nil)

	id = axsPolicy.PathParam(s.NewParam("ruleID", "Unique ID of access rule", s.P_Path, s.P_Required))
	id.Put("Update an access rule").
		Permission("applications::modify").
		Param(s.NewParam("access rule", "Access rule", s.P_Body, s.P_Schema(axsRule), s.P_Required)).
		Response(200, "OK", nil)
	id.Delete("Delete an access rule").
		Permission("applications::modify").
		Response(200, "OK", nil)

	id.PathParam(s.NewParam("index", "Desired index of rule", s.P_Path, s.P_Required, s.P_Number)).
		Patch("Move an access rule to a new place in the list").
		Permission("applications::modify").
		Response(200, "OK", nil)
}

const vnetModifyPerm = "virtual-networks::modify"
const vnetReadPerm = "virtual-networks::read"
const execReadPerm = "node-exec::read"
const execModifyPerm = "node-exec::modify"

func addV2DomainNetworkAPI(api *s.API) {
	p := api.Path("/v2/domain").PathParam(params.domainName).Path("/network")
	p.Produces("application/json").Consumes("application/json").Tag("Domain Network")

	vnet := api.Model("Virtual Network").
		Prop("name", s.NewSchema("Name of virtual network", s.S_String)).
		Prop("description", s.NewSchema("", s.S_String)).
		Prop("networkCidr", s.NewSchema("Virtual network CIDR", s.S_String, s.S_Example("10.10.14.0/24"))).
		Ref()

	p.Get("List domain networks").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(vnet))

	p.Post("Create a domain network").
		Permission(vnetModifyPerm).
		Param(s.NewParam("network", "Virtual network configuration", s.P_Body, s.P_Schema(vnet), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	p = p.PathParam(s.NewParam("networkName", "Network name", s.P_Path, s.P_Required))

	p.Delete("Delete a domain network - this change is not staged and will immediately affect the domain").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)

	accessPolicy := api.Model("Domain Network Access Policy").
		Prop("uid", s.NewSchema("", s.S_String)).
		Prop("protocol", s.NewSchema("Protocol", s.S_String, s.S_Enum("any", "tcp", "udp", "icmp"), s.S_Required)).
		Prop("source", s.NewSchema("", s.S_String, s.S_Required)).
		Prop("dest", s.NewSchema("", s.S_String, s.S_Required)).
		Prop("ports", s.NewSchema("", s.S_String)).
		Prop("notDest", s.NewSchema("", s.S_Boolean)).
		Prop("description", s.NewSchema("", s.S_String)).
		Ref()

	group := api.Model("Domain Network Group").
		Prop("name", s.NewSchema("Name", s.S_String)).
		Prop("description", s.NewSchema("", s.S_String)).
		Ref()

	groupMembership := api.Model("Domain Network Group Membership").
		Prop("objectName", s.NewSchema("", s.S_String)).
		Prop("groupName", s.NewSchema("", s.S_String)).
		Ref()

	authGroup := api.Model("Domain Network Authorization Group").
		Prop("name", s.NewSchema("Group name", s.S_String)).
		Prop("description", s.NewSchema("", s.S_String)).
		Ref()

	authGroupMembership := api.Model("Domain Network Authorization Group Membership").
		Prop("uid", s.NewSchema("", s.S_String)).
		Prop("name", s.NewSchema("", s.S_String)).
		Prop("memberType", s.NewSchema("", s.S_String)).
		Prop("idp", s.NewSchema("", s.S_String)).
		Prop("description", s.NewSchema("", s.S_String)).
		Ref()

	networkObject := api.Model("Domain Network Object").
		Prop("name", s.NewSchema("", s.S_String)).
		Prop("cidr", s.NewSchema("", s.S_String)).
		Prop("description", s.NewSchema("", s.S_String)).
		Ref()

	dnsConfig := api.Model("Domain Network DNS Configuration").
		Prop("enabled", s.NewSchema("", s.S_Boolean)).
		Prop("server", s.NewSchema("", s.S_String)).
		Ref()

	dnsZone := api.Model("Domain Network DNS Zone").
		Prop("name", s.NewSchema("", s.S_String)).
		Prop("resolver", s.NewSchema("", s.S_String)).
		Prop("description", s.NewSchema("", s.S_String)).
		Ref()

	dnsRecord := api.Model("Domain Network DNS Record").
		Prop("name", s.NewSchema("", s.S_String)).
		Prop("ttl", s.NewSchema("", s.S_Number)).
		Prop("recordType", s.NewSchema("", s.S_String)).
		Prop("value", s.NewSchema("", s.S_String)).
		Ref()

	route := api.Model("DomainNetworkRoute").
		Prop("uid", s.NewSchema("", s.S_String)).
		Prop("nodeName", s.NewSchema("", s.S_String)).
		Prop("networkCidr", s.NewSchema("", s.S_String)).
		Prop("metric", s.NewSchema("", s.S_Number)).
		Prop("description", s.NewSchema("", s.S_String)).
		Prop("domainName", s.NewSchema("", s.S_String)).
		Prop("networkName", s.NewSchema("", s.S_String)).
		Ref()

	inventory := api.Model("Domain Network Inventory").
		Prop("routes", s.NewArraySchema(route)).
		Prop("networkObjects", s.NewArraySchema(networkObject)).
		Prop("networkGroups", s.NewArraySchema(group)).
		Prop("networkGroupMemberships", s.NewArraySchema(groupMembership)).
		Prop("accessPolicies", s.NewArraySchema(accessPolicy)).
		Prop("dnsZones", s.NewArraySchema(dnsZone)).
		Prop("dnsRecords", s.NewArraySchema(dnsRecord)).
		Prop("dnsConfig", s.NewArraySchema(dnsConfig)).
		Ref()

	addAccessPolicyAPI(accessPolicy, p)
	addNetworkGroupAPI(group, groupMembership, p)
	addNetworkAuthGroupAPI(authGroup, authGroupMembership, p)
	addNetworkObjectAPI(networkObject, p)
	addNetworkDNSAPI(dnsConfig, dnsZone, dnsRecord, p)
	addNetworkRouteAPI(route, p)
	addDomainChangeManagementAPI(inventory, p)
}

func addDomainChangeManagementAPI(inventory *s.Schema, p *s.Path) {
	c := p.Path("/change")

	c.Get("List staged changes").
		Permission(vnetReadPerm).
		Response(200, "OK", inventory)

	digest := s.NewSchema("digest").
		Prop("digest", s.NewSchema("Digest of the network and its changes", s.S_String, s.S_Example("55ca6286e3e4f4fba5d0448333fa99fc5a404a73")))

	c.Path("/validate").Get("List validation errors for staged changes").
		Permission(vnetReadPerm).
		Response(200, "OK", digest).
		Response(422, "Validation failed", validationFailure)

	c.PathParam(s.NewParam("changeID", "Change ID", s.P_Path, s.P_Required)).
		Delete("Revert a staged change. If the item is newly added and not committed, the item will be deleted along with any associated changes.").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)

	c.Path("/commit").Post("Commit staged changes").
		Permission(vnetModifyPerm).
		Param(s.NewParam("digest", "Digest", s.P_Body, s.P_Schema(digest), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
}

func addNetworkRouteAPI(route *s.Schema, p *s.Path) {
	r := p.Path("/route")
	r.Get("List a network's routes").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(route))
	r.Post("Create a network route").
		Permission(vnetModifyPerm).
		Param(s.NewParam("route", "Route configuration", s.P_Body, s.P_Schema(route), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	r = r.PathParam(s.NewParam("routeID", "Route ID", s.P_Path, s.P_Required))
	r.Put("Update a network route").
		Permission(vnetModifyPerm).
		Param(s.NewParam("route", "Route configuration", s.P_Body, s.P_Schema(route), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	r.Delete("Delete a network route").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)
}

func addNetworkDNSAPI(dnsConfig, dnsZone, dnsRecord *s.Schema, p *s.Path) {
	dns := p.Path("/dns")
	dns.Get("Get a network's DNS configuration").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(dnsConfig))
	dns.Put("Update a network's DNS configuration").
		Permission(vnetModifyPerm).
		Param(s.NewParam("dns", "DNS configuration", s.P_Body, s.P_Schema(dnsConfig), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	dns = p.Path("/dns-zone")
	dns.Get("List a network's DNS zones").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(dnsZone))
	dns.Post("Create a DNS zone").
		Permission(vnetModifyPerm).
		Param(s.NewParam("zone", "DNS zone configuration", s.P_Body, s.P_Schema(dnsZone), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	dns = dns.PathParam(s.NewParam("zoneName", "DNS zone name", s.P_Path, s.P_Required))
	dns.Put("Update a DNS zone").
		Permission(vnetModifyPerm).
		Param(s.NewParam("zone", "DNS zone configuration", s.P_Body, s.P_Schema(dnsZone), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	dns.Delete("Delete a DNS zone").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)

	dns = dns.Path("/dns-record")
	dns.Get("List a network zone's DNS records").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(dnsRecord))
	dns.Post("Create a DNS record").
		Permission(vnetModifyPerm).
		Param(s.NewParam("record", "DNS record configuration", s.P_Body, s.P_Schema(dnsRecord), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	dns = dns.
		PathParam(s.NewParam("recordName", "DNS record name", s.P_Path, s.P_Required)).
		PathParam(s.NewParam("recordType", "DNS record type", s.P_Path, s.P_Required))

	dns.Put("Update a DNS record").
		Permission(vnetModifyPerm).
		Param(s.NewParam("record", "DNS record configuration", s.P_Body, s.P_Schema(dnsRecord), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	dns.Delete("Delete a DNS record").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)
}

func addNetworkObjectAPI(networkObject *s.Schema, p *s.Path) {
	no := p.Path("/network-object")
	no.Get("List a network's objects").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(networkObject))
	no.Post("Create a network object").
		Permission(vnetModifyPerm).
		Param(s.NewParam("object", "Network object configuration", s.P_Body, s.P_Schema(networkObject), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	no = no.PathParam(s.NewParam("objectName", "Network object name", s.P_Path, s.P_Required))
	no.Put("Update a network object").
		Permission(vnetModifyPerm).
		Param(s.NewParam("object", "Network object configuration", s.P_Body, s.P_Schema(networkObject), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	no.Delete("Delete a network object").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)
}

func addNetworkAuthGroupAPI(authGroup *s.Schema, authGroupMembership *s.Schema, p *s.Path) {
	ag := p.Path("/auth-group")
	ag.Get("List a network's auth groups").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(authGroup))
	ag.Post("Create a network auth group").
		Permission(vnetModifyPerm).
		Param(s.NewParam("authGroup", "Network auth group configuration", s.P_Body, s.P_Schema(authGroup), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	ag = ag.PathParam(s.NewParam("groupName", "Network auth group name", s.P_Path, s.P_Required))
	ag.Get("List a network auth group's members").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(authGroupMembership))
	ag.Delete("Delete a network auth group").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)
	ag.Put("Update a network auth group").
		Permission(vnetModifyPerm).
		Param(s.NewParam("authGroup", "Network auth group configuration", s.P_Body, s.P_Schema(authGroup), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	ag.Post("Add a network auth group member").
		Permission(vnetModifyPerm).
		Param(s.NewParam("member", "Network auth group member", s.P_Body, s.P_Schema(authGroupMembership), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	ag = ag.PathParam(s.NewParam("memberID", "Group member ID", s.P_Path, s.P_Required))
	ag.Delete("Remove a network auth group member").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)
	ag.Put("Update a membership").
		Permission(vnetModifyPerm).
		Param(s.NewParam("member", "Network auth group member", s.P_Body, s.P_Schema(authGroupMembership), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
}

func addNetworkGroupAPI(group *s.Schema, groupMembership *s.Schema, p *s.Path) {
	nwg := p.Path("/network-group")
	nwg.Get("List a network's groups").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(group))

	nwg.Post("Create a network group").
		Permission(vnetModifyPerm).
		Param(s.NewParam("group", "Network group configuration", s.P_Body, s.P_Schema(group), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	nwg = nwg.PathParam(s.NewParam("groupName", "Network group name", s.P_Path, s.P_Required))
	nwg.Get("List a network's group memberships").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(groupMembership))
	nwg.Put("Update a network group").
		Permission(vnetModifyPerm).
		Param(s.NewParam("group", "Network group configuration", s.P_Body, s.P_Schema(group), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	nwg.Delete("Delete a network group").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)

	nwg = nwg.PathParam(s.NewParam("objectName", "Network object name", s.P_Path, s.P_Required))
	nwg.Delete("Remove a network object from a network group").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)
	nwg.Post("Add a network object to a network group (represented by a network group membership").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil)
}

func addAccessPolicyAPI(accessPolicy *s.Schema, p *s.Path) {
	axs := p.Path("/access-policy")

	axs.Get("List a network's access policies").
		Permission(vnetReadPerm).
		Response(200, "OK", s.NewArraySchema(accessPolicy))

	axs.Post("Create a network access policy").
		Permission(vnetModifyPerm).
		Param(s.NewParam("acl", "Access policy configuration", s.P_Body, s.P_Schema(accessPolicy), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	axs = axs.PathParam(s.NewParam("accessPolicyID", "Unique ID of access policy", s.P_Path, s.P_Required))
	axs.Put("Update a network access policy").
		Permission(vnetModifyPerm).
		Param(s.NewParam("acl", "Access policy configuration", s.P_Body, s.P_Schema(accessPolicy), s.P_Required)).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	axs.Delete("Delete a network access policy").
		Permission(vnetModifyPerm).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
}

func addSwaggerYML(api *s.API) {
	api.AddSwaggerYAML("swagger.yml")
}

func addAuditAPI(api *s.API) {
	p := api.Path("/audit").Consumes("application/json").Tag("Audit")

	configAudit := s.NewSchema("ConfigChange").
		Prop("uid", s.NewSchema("Unique ID of the change", s.S_Example(uuid.NewString()))).
		Prop("timestamp", s.NewSchema("Unix timestamp when the change happened", s.S_Example(time.Now().Unix()))).
		Prop("ip", s.NewSchema("IP address of the client that initiated the change", s.S_Example("44.44.44.44"))).
		Prop("userName", s.NewSchema("User name of the client that initiated the change", s.S_Example("admin@trustgrid.io"))).
		Prop("itemType", s.NewSchema("Type of the item that was changed", s.S_Example("Node"))).
		Prop("itemId", s.NewSchema("Unique ID of the item that was changed", s.S_Example(examples.nodeID))).
		Prop("auditType", s.NewSchema("Type of change", s.S_Enum("change", "create", "delete", "action"))).
		Prop("message", s.NewSchema("Message describing the change", s.S_Example("Node license created")))

	p.Path("/tail/config").
		Permission("audits::read:config").
		Produces("application/json").
		Get("List configuration change audits").
		Param(s.NewParam("timestamp", "Cursor returned from the previous request", s.P_Query)).
		Response(200, "OK", s.NewArraySchema(configAudit))

	p.Path("/download/config").
		Permission("audits::read:config").
		Produces("text/csv").
		Get("Download configuration change audits").
		Param(s.NewParam("itemID", "ID for the item to audit. If specified, must include itemType.", s.P_Query)).
		Param(s.NewParam("itemType", "Type of item to audit, e.g., Node or Cluster. If specified, must include itemID", s.P_Query)).
		Response(200, "OK", nil)

	tcpFlags := []string{"* 0x01 - FIN", "* 0x02 - SYN", "* 0x04 - RST", "* 0x08 - PSH", "* 0x10 - ACK", "* 0x20 - URG"}

	flowLog := s.NewSchema("FlowLog").
		Prop("destBytes", s.NewSchema("Destination bytes transferred", s.S_Example(100), s.S_Number)).
		Prop("sourceBytes", s.NewSchema("Source bytes transferred", s.S_Example(100), s.S_Number)).
		Prop("destIP", s.NewSchema("Destination IP address", s.S_Example("1.1.1.1"))).
		Prop("destPort", s.NewSchema("Destination port", s.S_Example(1234), s.S_Number)).
		Prop("sourceIP", s.NewSchema("Source IP address", s.S_Example("2.2.2.2"))).
		Prop("sourcePort", s.NewSchema("Source port", s.S_Example(1234), s.S_Number)).
		Prop("destNode", s.NewSchema("Destination node name", s.S_Example(examples.nodeName))).
		Prop("sourceNode", s.NewSchema("Source node name", s.S_Example(examples.nodeName))).
		Prop("startTime", s.NewSchema("Start time of the flow", s.S_Example("2022-11-01T22:46:01.765Z"))).
		Prop("endTime", s.NewSchema("Start time of the flow", s.S_Example("2022-11-01T22:46:02.730Z"))).
		Prop("ztnaAppID", s.NewSchema("ZTNA application ID", s.S_Example(examples.appID))).
		Prop("ztnaSessionID", s.NewSchema("ZTNA application ID", s.S_Example(examples.appID))).
		Prop("activityID", s.NewSchema("Threat Intelligence activity ID", s.S_Example(examples.appID))).
		Prop("protocol", s.NewSchema("Protocol", s.S_Example("TCP"), s.S_Enum("TCP", "UDP", "ICMP", "UNKNOWN"))).
		Prop("tcpFlags", s.NewSchema("TCP Flags encoded in hex:\n"+strings.Join(tcpFlags, "\n"), s.S_Example("00100001")))

	tcpFlagParam := s.NewArraySchema(s.NewSchema("tcpFlags", s.S_Number))
	tcpFlagParam.Example = []int{1, 2}

	p.Path("/tail/flow_logs").
		Permission("audits::read:flows").
		Produces("application/json").
		Get("List recent flow logs").
		Param(s.NewParam("sTime", "Unix timestamp for the start of the log window", s.P_Query, s.P_Number)).
		Param(s.NewParam("eTime", "Unix timestamp for the end of the log window", s.P_Query, s.P_Number)).
		Param(s.NewParam("eTimeOp", "Comparison operator for the end of the log window", s.P_Query, s.P_Enum("eq", "ne", "gt", "gte", "lt", "lte"))).
		Param(s.NewParam("protocol", "IP protocol", s.P_Query)).
		Param(s.NewParam("srcIp", "Source IP address", s.P_Query)).
		Param(s.NewParam("dstIp", "Destination IP address", s.P_Query)).
		Param(s.NewParam("srcPort", "Source port", s.P_Query, s.P_Number)).
		Param(s.NewParam("srcPortOp", "Comparison operator for the source port", s.P_Query, s.P_Enum("eq", "ne", "gt", "gte", "lt", "lte"))).
		Param(s.NewParam("dstPort", "Destination port", s.P_Query, s.P_Number)).
		Param(s.NewParam("dstPortOp", "Comparison operator for the dest port", s.P_Query, s.P_Enum("eq", "ne", "gt", "gte", "lt", "lte"))).
		Param(s.NewParam("limit", "Maximum number of results to return", s.P_Query, s.P_Number)).
		Param(s.NewParam("srcNode", "Source node name", s.P_Query)).
		Param(s.NewParam("dstNode", "Dest node name", s.P_Query)).
		Param(s.NewParam("node", "Flow logging node ID", s.P_Query)).
		Param(s.NewParam("reverse", "When true, newer flow logs will be listed first", s.P_Query, s.P_Boolean)).
		Param(s.NewParam("tcpFlags", "If provided, a flow must match at least one of the TCP flags provided. Decimal encoded, see flow log TCP flag encoding.", s.P_Query, s.P_Array, s.P_Schema(tcpFlagParam))).
		Param(s.NewParam("cursor", "Continuation cursor from previous query", s.P_Query)).
		Response(200, "OK", s.NewArraySchema(flowLog), s.Header{Name: "x-cursor", Type: "string", Description: "Continuation cursor for the next query"})

	// TODO add output here and params if we support them?
	p.Path("/tail/user").
		Permission("audits::read:user").
		Produces("application/json").
		Get("List authentication audits").
		Response(200, "OK", nil)

	p.Path("/download/user").
		Permission("audits::read:user").
		Produces("text/csv").
		Get("Download authentication audits").
		Response(200, "OK", nil)

		// TODO add output type here
	p.Path("/tail/node").
		Permission("audits::read:node").
		Produces("application/json").
		Get("List node audits").
		Param(s.NewParam("timestamp", "Cursor returned from the previous request", s.P_Query)).
		Param(s.NewParam("FQDN", "Node FQDN", s.P_Query)).
		Response(200, "OK", nil)

	p.Path("/download/node").
		Permission("audits::read:node").
		Produces("text/csv").
		Get("Download node audits").
		Param(s.NewParam("timestamp", "Cursor returned from the previous request", s.P_Query)).
		Param(s.NewParam("FQDN", "Node FQDN", s.P_Query)).
		Response(200, "OK", nil)
}

func addNodeAPI(api *s.API) {
	p := api.Path("/node").Produces("application/json").Consumes("application/json").Tag("Node")

	node := api.Model("Node").
		Prop("uid", s.NewSchema("Node ID", s.S_String, s.S_Example(examples.nodeID))).
		Prop("name", s.NewSchema("Node name", s.S_String, s.S_Example("mynode"))).
		Prop("state", s.NewSchema("Node state", s.S_String, s.S_Enum("ACTIVE", "INACTIVE"))).
		Prop("cluster", s.NewSchema("Cluster FQDN", s.S_String, s.S_Example(examples.clusterFQDN))).
		Prop("config", s.NewSchema("Node config")).
		Prop("shadow", s.NewSchema("Node shadow")).
		Prop("tags", s.NewSchema("Node tags")).
		Prop("online", s.NewSchema("True when the node is connected to the control plane", s.S_Boolean)).
		Ref()

	projection := s.NewArraySchema(s.NewSchema("projection", s.S_String))
	projection.Example = []string{"uid", `["config", "apigw", "enabled"]`}

	p.Get("List nodes").
		Permission("nodes::read").
		Param(s.NewParam("tags", "Comma-separated key:value pairs for tag filtering, e.g., location:Austin,device:Trustgrid", s.P_Query)).
		Param(s.NewParam("projection", "List of fields to return from the API. Supports nested fields and anything in the Node schema", s.P_Query, s.P_Array, s.P_Schema(projection))).
		Response(200, "OK", s.NewArraySchema(node))

	p = p.PathParam(params.nodeID)

	p.Get("Get a node").
		Permission("nodes::read").
		Response(200, "OK", node).
		Response(404, "Not Found", nil)

	p.Path("/alerts")
	updates := s.NewSchema("Node updates").
		Prop("state", s.NewSchema("Node state", s.S_String, s.S_Enum("ACTIVE", "INACTIVE"))).
		Prop("cluster", s.NewSchema("Cluster FQDN - requires `nodes::cluster` permission to modify", s.S_String, s.S_Example(examples.clusterFQDN)))

	p.Put("Update a node").
		Permission("nodes::manage").
		Param(s.NewParam("updates", "Node updates", s.P_Body, s.P_Schema(updates))).
		Response(200, "OK", nil)

	p.Delete("Delete a node").
		Permission("nodes::delete").
		Response(200, "OK", nil).
		Response(404, "Not Found", nil)
}

func addNodeDatastoreAPI(api *s.API) {
	p := api.Path("/v2/node").PathParam(params.nodeID).Path("/data-store").Produces("application/json").Consumes("application/json").Tag("Node").Permission("nodes::service:datastore-manager")

	file := api.Model("Data Store File").
		Prop("path", s.NewSchema("File path", s.S_String, s.S_Example("qcows/win2019-serverli.qcow2"))).
		Prop("size", s.NewSchema("File size", s.S_String, s.S_Example("9.4 GB"))).
		Prop("name", s.NewSchema("File name", s.S_String, s.S_Example("win2019-serverli.qcow2"))).
		Prop("type", s.NewSchema("File type", s.S_String, s.S_Enum("file", "directory"))).
		Prop("files", s.NewSchema("Children, if a directory. This is a recursive model.", s.S_String)).
		Ref()

	fileList := s.NewSchema("File List").
		Prop("files", s.NewArraySchema(file))

	p.Path("/list").Get("List the data store contents").
		Response(200, "OK", fileList)

	mkdir := s.NewSchema("mkdir").
		Prop("filename", s.NewSchema("filename", s.S_String, s.S_Required))

	p.Path("/mkdir").Post("Create a directory").
		Param(s.NewParam("mkdir", "mkdir", s.P_Body, s.P_Schema(mkdir))).
		Response(200, "OK", nil)

	p.Delete("Delete a file or directory").
		Param(s.NewParam("rm", "rm", s.P_Body, s.P_Schema(mkdir))).
		Response(200, "OK", nil)

	httpUpload := s.NewSchema("Data Store HTTP Upload Request").
		Prop("uri", s.NewSchema("Destination URI", s.S_String, s.S_Example("http://example.com/file.iso"))).
		Prop("filePath", s.NewSchema("File path in data store", s.S_String, s.S_Example("file.iso"))).
		Prop("multipart", s.NewSchema("Multipart upload", s.S_Boolean, s.S_Example(false)))

	p.Path("/http-upload").
		Post("Upload a file from the node to an HTTP endpoint").
		Param(s.NewParam("config", "Upload config", s.P_Body, s.P_Schema(httpUpload))).
		Response(200, "OK", nil)

	s3Upload := s.NewSchema("Data Store S3 Upload Request").
		Prop("bucketDest", s.NewSchema("S3 destination", s.S_String, s.S_Example("s3://your-bucket/your-destination.iso"))).
		Prop("filePath", s.NewSchema("File path in data store", s.S_String, s.S_Example("file.iso")))

	p.Path("/s3-upload").
		Post("Upload a file from the node to an S3 bucket").
		Param(s.NewParam("config", "Upload config", s.P_Body, s.P_Schema(s3Upload))).
		Response(200, "OK", nil)

	s3Download := s.NewSchema("Data Store S3 Download Request").
		Prop("uri", s.NewSchema("S3 location of source file", s.S_String, s.S_Example("s3://your-bucket/your-file.iso"))).
		Prop("location", s.NewSchema("Target destination in data store", s.S_String, s.S_Example("your-file.iso")))

	p.Path("/s3-download").
		Post("Copy a file from S3 to the data store").
		Param(s.NewParam("config", "Download config", s.P_Body, s.P_Schema(s3Download))).
		Response(200, "OK", nil)

	httpDownload := s.NewSchema("Data Store HTTP Download Request").
		Prop("uri", s.NewSchema("URI of source file", s.S_String, s.S_Example("http://example.com/file.iso"))).
		Prop("filename", s.NewSchema("Target filename", s.S_String, s.S_Example("file.iso"))).
		Prop("location", s.NewSchema("Target destination directory in data store", s.S_String, s.S_Example("/isos"))).
		Prop("hash", s.NewSchema("Hash of the file", s.S_String, s.S_Example("07c43d77ab8d30cee094ce09b14f87fa"))).
		Prop("algo", s.NewSchema("Hashing algorithm of the file - MD5, SHA1, or SHA-256", s.S_String, s.S_Example("MD5")))

	p.Path("/http-download").
		Post("Copy a file from an HTTP location to the data store").
		Param(s.NewParam("config", "Download config", s.P_Body, s.P_Schema(httpDownload))).
		Response(200, "OK", nil)

	task := api.Model("DataStoreTask").
		Prop("eTime", s.NewSchema("End timestamp", s.S_Number, s.S_Example(1667861503331))).
		Prop("name", s.NewSchema("Task name", s.S_String, s.S_Example("Upload S3 File"))).
		Prop("details", s.NewSchema("Task details", s.S_String, s.S_Example("Finished the file transfer"))).
		Prop("id", s.NewSchema("ID", s.S_String, s.S_Example("0d7f7145-32bd-45e9-ac4b-9cf04887392"))).
		Prop("sTime", s.NewSchema("Start timestamp", s.S_Number, s.S_Example(1667861503331))).
		Prop("status", s.NewSchema("Task status", s.S_String, s.S_Example("Completed"))).
		Ref()

	taskList := s.NewSchema("Task List").
		Prop("tasks", s.NewArraySchema(task))

	p.Path("/tasks").Get("List recent data store activity").Response(200, "OK", taskList)
}

func addNodeConfigAPIs(api *s.API) {
	p := api.Path("/node").PathParam(params.nodeID).Produces("application/json").Consumes("application/json").Tag("Node")

	gatewayConfig := s.NewSchema("Gateway config").
		Prop("enabled", s.NewSchema("Enable gateway plugin", s.S_Boolean, s.S_Required)).
		Prop("host", s.NewSchema("Hostname of the gateway", s.S_String, s.S_Example("mygateway.trustgrid.io"))).
		Prop("port", s.NewSchema("Port of the gateway", s.S_Number, s.S_Example(8080))).
		Prop("maxmbps", s.NewSchema("Max ingress MBPS", s.S_Number, s.S_Example(1000))).
		Prop("cert", s.NewSchema("Certificate", s.S_String, s.S_Example("mygateway.trustgrid.io"))).
		Prop("type", s.NewSchema("Type of gateway", s.S_String, s.S_Enum("private", "public", "hub"))).
		Prop("connectToPublic", s.NewSchema("Connect to public", s.S_Boolean)).
		Prop("maxClientWriteMbps", s.NewSchema("Max egress MBPS", s.S_Number, s.S_Example(1000))).
		Prop("udpEnabled", s.NewSchema("Enable UDP", s.S_Boolean)).
		Prop("monitorHops", s.NewSchema("Monitor hops", s.S_Boolean)).
		Prop("udpPort", s.NewSchema("UDP port", s.S_Number, s.S_Example(8081)))

	p.Path("/gateway").Put("Update gateway configuration").
		Permission("nodes::configure:gateway").
		Param(s.NewParam("config", "Gateway config", s.P_Body, s.P_Schema(gatewayConfig))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
}

func addExecAPI(api *s.API) {
	node := api.Path("/v2/node").PathParam(params.nodeID).Path("/exec").Produces("application/json").Consumes("application/json").Tag("Compute")
	cluster := api.Path("/v2/cluster").PathParam(params.clusterFQDN).Path("/exec").Produces("application/json").Consumes("application/json").Tag("Compute")

	image := api.Model("Image").
		Prop("repository", s.NewSchema("Image repository", s.S_String, s.S_Example("mycompany.trustgrid.io/myimage"))).
		Prop("tag", s.NewSchema("Image tag", s.S_String, s.S_Example("latest"))).
		Ref()

	container := api.Model("Container").
		Prop("id", s.NewSchema("Container ID", s.S_String, s.S_Example(uuid.NewString()))).
		Prop("name", s.NewSchema("Container name", s.S_String, s.S_Example("mycontainer"))).
		Prop("enabled", s.NewSchema("", s.S_Boolean)).
		Prop("execType", s.NewSchema("Execution type", s.S_String, s.S_Enum("onDemand", "service", "recurring"))).
		Prop("description", s.NewSchema("Container description", s.S_String, s.S_Example("my container"))).
		Prop("image", image).
		Prop("hostname", s.NewSchema("Container hostname", s.S_String, s.S_Example("mycontainer"))).
		Prop("stopTime", s.NewSchema("Grace period for container to stop when requested, in seconds", s.S_String, s.S_Example("60"))).
		Prop("useInit", s.NewSchema("Indicates that an init process should be used as PID 1 in the container. Ensures responsibilities of an init system are performed inside the container (e.g., handling exit signals).", s.S_Boolean)).
		Prop("privileged", s.NewSchema("Run the container as a privileged user", s.S_Boolean)).
		Prop("requireConnectivity", s.NewSchema("Only start the container if the node has connectivity to the control plane. Needed for encrypted volumes.", s.S_Boolean)).
		Prop("command", s.NewSchema("Command to run in the container", s.S_String)).
		Prop("user", s.NewSchema("User name (or UID) and optionally the group (or GID) to use when starting the container.  This will override the USER specified in the image", s.S_String)).
		Ref()

	n := node.Path("/container")
	c := cluster.Path("/container")

	n.Get("List containers").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(container))
	c.Get("List containers").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(container))

	n.Post("Create a container").
		Permission(execModifyPerm).
		Param(s.NewParam("container", "Container", s.P_Body, s.P_Schema(container))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	c.Post("Create a container").
		Permission(execModifyPerm).
		Param(s.NewParam("container", "Container", s.P_Body, s.P_Schema(container))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	n = n.PathParam(s.NewParam("containerID", "Container ID", s.P_Path, s.P_Required))
	c = c.PathParam(s.NewParam("containerID", "Container ID", s.P_Path, s.P_Required))
	n.Get("Get a container").
		Permission(execReadPerm).
		Response(200, "OK", container)
	n.Delete("Delete a container").
		Permission(execModifyPerm).
		Response(200, "OK", nil)
	n.Put("Update a container").
		Permission(execModifyPerm).
		Param(s.NewParam("container", "Container", s.P_Body, s.P_Schema(container))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	c.Get("Get a container").
		Permission(execReadPerm).
		Response(200, "OK", container)
	c.Delete("Delete a container").
		Permission(execModifyPerm).
		Response(200, "OK", nil)
	c.Put("Update a container").
		Permission(execModifyPerm).
		Param(s.NewParam("container", "Container", s.P_Body, s.P_Schema(container))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	variable := api.Model("Container Variable").
		Prop("name", s.NewSchema("Variable name", s.S_String, s.S_Example("myvar"))).
		Prop("value", s.NewSchema("Variable value", s.S_String, s.S_Example("myvalue"))).
		Ref()

	mount := api.Model("Container Mount").
		Prop("uid", s.NewSchema("UID", s.S_String, s.S_Example(uuid.NewString()))).
		Prop("mountType", s.NewSchema("Mount type", s.S_String, s.S_Enum("volume", "bind"))).
		Prop("source", s.NewSchema("Source", s.S_String, s.S_Example("myvolume"))).
		Prop("dest", s.NewSchema("Destination", s.S_String, s.S_Example("/myvolume"))).
		Prop("encrypted", s.NewSchema("Encrypted", s.S_Boolean)).
		Ref()

	portMapping := api.Model("Container Port Mapping").
		Prop("uid", s.NewSchema("UID", s.S_String, s.S_Example(uuid.NewString()))).
		Prop("protocol", s.NewSchema("Protocol", s.S_String, s.S_Enum("tcp", "udp"))).
		Prop("iface", s.NewSchema("Interface", s.S_String, s.S_Example("eth0"))).
		Prop("hostPort", s.NewSchema("Host port", s.S_String, s.S_Example("8080"))).
		Prop("containerPort", s.NewSchema("Container port", s.S_String, s.S_Example("8080"))).
		Ref()

	virtualNetwork := api.Model("Container Virtual Network").
		Prop("uid", s.NewSchema("UID", s.S_String, s.S_Example(uuid.NewString()))).
		Prop("network", s.NewSchema("Network name", s.S_String, s.S_Example("mynetwork"))).
		Prop("ip", s.NewSchema("Virtual IP", s.S_String, s.S_Example("10.10.10.14"))).
		Prop("allowOutbound", s.NewSchema("Allow outbound traffic", s.S_Boolean)).
		Ref()

	limit := api.Model("Container Limits").
		Prop("cpuMax", s.NewSchema("CPU max", s.S_Number, s.S_Example(80))).
		Prop("memMax", s.NewSchema("Memory max", s.S_Number, s.S_Example(80))).
		Prop("legacyMemMax", s.NewSchema("Legacy memory max", s.S_Number, s.S_Example(80))).
		Prop("memHigh", s.NewSchema("Memory high", s.S_Number, s.S_Example(80))).
		Prop("legacyMemHigh", s.NewSchema("Legacy memory high", s.S_Number, s.S_Example(80))).
		Prop("ioRbps", s.NewSchema("IO read rate", s.S_Number, s.S_Example(80))).
		Prop("ioWbps", s.NewSchema("IO write rate", s.S_Number, s.S_Example(80))).
		Prop("ioRiops", s.NewSchema("IO read IOPS", s.S_Number, s.S_Example(80))).
		Prop("ioWiops", s.NewSchema("IO write IOPS", s.S_Number, s.S_Example(80))).
		Prop("limits", s.NewArraySchema(s.NewSchema("Limit").
			Prop("type", s.NewSchema("Limit type", s.S_String, s.S_Enum("core", "cpu", "data", "fsize", "locks", "memlock", "msgqueue", "nice", "nofile", "nproc", "rss", "rtprio", "rttime", "sigpending", "stack"))).
			Prop("soft", s.NewSchema("Soft limit", s.S_Number, s.S_Example(10))).
			Prop("hard", s.NewSchema("Hard limit", s.S_Number, s.S_Example(20))))).
		Ref()

	healthCheck := api.Model("Container Health Check").
		Prop("command", s.NewSchema("Command", s.S_String, s.S_Example("/bin/sh -c 'exit 0'"))).
		Prop("interval", s.NewSchema("Interval", s.S_Number, s.S_Example(10))).
		Prop("timeout", s.NewSchema("Timeout", s.S_Number, s.S_Example(10))).
		Prop("startPeriod", s.NewSchema("Grace period before health checks are monitored, in seconds", s.S_Number, s.S_Example(10))).
		Prop("retries", s.NewSchema("Number of health check retries before a container is marked unhealthy", s.S_Number, s.S_Example(10), s.S_Default(3))).
		Ref()

	caps := api.Model("Container Capabilities").
		Prop("addCaps", s.NewArraySchema(s.NewSchema("Added capabilities", s.S_String))).
		Prop("dropCaps", s.NewArraySchema(s.NewSchema("Dropped capabilities", s.S_String))).
		Ref()

	logging := api.Model("Container Logging").
		Prop("maxFileSize", s.NewSchema("Maximum file size in MB", s.S_Number, s.S_Example(10))).
		Prop("numFiles", s.NewSchema("Maximum number of files", s.S_Number, s.S_Example(10))).
		Ref()

	vrf := api.Model("Container VRF").
		Prop("name", s.NewSchema("VRF name", s.S_String, s.S_Example("myvrf"))).
		Ref()

	iface := api.Model("Container Interfaces").
		Prop("uid", s.NewSchema("UID", s.S_String, s.S_Example(uuid.NewString()))).
		Prop("name", s.NewSchema("Interface name", s.S_String, s.S_Example("eth0"))).
		Prop("dest", s.NewSchema("Destination", s.S_String, s.S_Example("10.10.14.0/24"))).
		Ref()

	config := api.Model("ContainerConfig").
		Prop("variables", s.NewArraySchema(variable)).
		Prop("mounts", s.NewArraySchema(mount)).
		Prop("portMappings", s.NewArraySchema(portMapping)).
		Prop("virtualNetworks", s.NewArraySchema(virtualNetwork)).
		Prop("limits", limit).
		Prop("healthCheck", healthCheck).
		Prop("capabilities", caps).
		Prop("logging", logging).
		Prop("vrf", vrf).
		Prop("interfaces", s.NewArraySchema(iface)).
		Ref()

	volume := api.Model("Volume").
		Prop("name", s.NewSchema("Volume name", s.S_String, s.S_Example("myvolume"))).
		Prop("encrypted", s.NewSchema("Encrypted", s.S_Boolean)).
		Ref()

	n.Path("/config").Put("Update container config").
		Permission(execModifyPerm).
		Param(s.NewParam("config", "Container config", s.P_Body, s.P_Schema(config))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	c.Path("/config").Put("Update container config").
		Permission(execModifyPerm).
		Param(s.NewParam("config", "Container config", s.P_Body, s.P_Schema(config))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	n.Path("/variable").Get("List container variables").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(variable))
	c.Path("/variable").Get("List container variables").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(variable))

	n.Path("/port-mapping").Get("List container port mappings").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(portMapping))
	c.Path("/port-mapping").Get("List container port mappings").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(portMapping))

	n.Path("/virtual-network").Get("List container virtual networks").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(virtualNetwork))
	c.Path("/virtual-network").Get("List container virtual networks").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(virtualNetwork))

	n.Path("/interface").Get("List container interfaces").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(iface))
	c.Path("/interface").Get("List container interfaces").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(iface))

	n.Path("/vrf").Get("Get container VRF").
		Permission(execReadPerm).
		Response(200, "OK", vrf)
	c.Path("/vrf").Get("Get container VRF").
		Permission(execReadPerm).
		Response(200, "OK", vrf)

	n.Path("/mount").Get("List container mounts").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(mount))
	c.Path("/mount").Get("List container mounts").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(mount))

	n.Path("/volume").Get("List container volumes").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(volume))
	c.Path("/volume").Get("List container volumes").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(volume))

	n.Path("/limit").Get("List container limits").
		Permission(execReadPerm).
		Response(200, "OK", limit)
	c.Path("/limit").Get("List container limits").
		Permission(execReadPerm).
		Response(200, "OK", limit)

	n.Path("/healthcheck").Get("Get container health check").
		Permission(execReadPerm).
		Response(200, "OK", healthCheck)
	c.Path("/healthcheck").Get("Get container health check").
		Permission(execReadPerm).
		Response(200, "OK", healthCheck)

	n.Path("/capability").Get("Get container capabilities").
		Permission(execReadPerm).
		Response(200, "OK", caps)
	c.Path("/capability").Get("Get container capabilities").
		Permission(execReadPerm).
		Response(200, "OK", caps)

	n.Path("/logging").Get("Get container logging configuration").
		Permission(execReadPerm).
		Response(200, "OK", logging)
	c.Path("/logging").Get("Get container logging configuration").
		Permission(execReadPerm).
		Response(200, "OK", logging)

	n = node.Path("/volume")
	c = cluster.Path("/volume")

	n.Get("List volumes").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(volume))
	c.Get("List volumes").
		Permission(execReadPerm).
		Response(200, "OK", s.NewArraySchema(volume))

	n.Post("Create volume").
		Permission(execModifyPerm).
		Param(s.NewParam("volume", "Volume", s.P_Body, s.P_Schema(volume))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)
	c.Post("Create volume").
		Permission(execModifyPerm).
		Param(s.NewParam("volume", "Volume", s.P_Body, s.P_Schema(volume))).
		Response(200, "OK", nil).
		Response(422, "Validation failed", validationFailure)

	n.PathParam(s.NewParam("volumeName", "Volume Name", s.P_Path, s.P_Required)).Delete("Delete a volume").
		Permission(execModifyPerm).
		Response(200, "OK", nil)
	c.PathParam(s.NewParam("volumeName", "Volume Name", s.P_Path, s.P_Required)).Delete("Delete a volume").
		Permission(execModifyPerm).
		Response(200, "OK", nil)
}

func main() {
	params.domainName = s.NewParam("domainName", "Domain name", s.P_Path, s.P_Required)
	params.nodeID = s.NewParam("nodeID", "Node ID", s.P_Path, s.P_Required)
	params.clusterFQDN = s.NewParam("clusterFQDN", "Node FQDN", s.P_Path, s.P_Required)

	examples.nodeID = uuid.NewString()
	examples.appID = uuid.NewString()
	examples.sessionID = uuid.NewString()
	examples.nodeName = "mynode"
	examples.fqdn = "mynode.trustgrid.io"
	examples.clusterFQDN = "mycluster.trustgrid.io"

	api := s.NewAPI()
	api.Info.Title = "Trustgrid Management API"
	api.Info.Version = "1.0.0"
	api.Schemes = []string{"https"}
	api.Host = "api.trustgrid.io"
	api.Swagger = "2.0"
	validationFailure = api.Model("ValidationFailed")
	validationFailure.Type = s.ST_Array
	validationFailure.Items = s.NewSchema("Validation failure", s.S_String)

	addAlertAPI(api)
	addV2AlertAPI(api)
	addTagsAPI(api)
	addCertificatesAPI(api)
	addDomainAPI(api)
	addApplicationAPI(api)
	addV2DomainNetworkAPI(api)
	addAuditAPI(api)
	addNodeAPI(api)
	addNodeConfigAPIs(api)
	addNodeDatastoreAPI(api)
	addExecAPI(api)
	addSwaggerYML(api)

	out, err := json.MarshalIndent(api, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
