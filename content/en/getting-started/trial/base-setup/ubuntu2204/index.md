---
title: Deploying Ubuntu VMs
linkTitle: Ubuntu VMs
description: Details methods for deploying Ubuntu instances required for the trial steps
---

The trial requires users have access to two instances of Ubuntu 22.04 LTS (Jammy Jellyfish) to on which the Trustgrid agent is installed. There are numerous ways to acheive this, but below are a few methods to streamline the process.


## Deploy on AWS
The below methods assume you have access to an AWS account. The EC2 resources required fall under the [AWS Free Tier](https://aws.amazon.com/free) program so costs should be minimal.

### AWS via Terraform

Users familiar with [Hashicorp Terraform](https://developer.hashicorp.com/terraform) can spin up EC2 instances in multiple regions using the below Terraform example.  

#### Requirements
- [Terraform](https://developer.hashicorp.com/terraform/install) installed
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) with a [profile configured](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html#getting-started-quickstart-new-command).
    - The below process assumes you name this profile `default`
- [SSH key pair](https://www.digitalocean.com/community/tutorials/how-to-create-ssh-keys-with-openssh-on-macos-or-linux#step-3-generating-keys-with-openssh) with both private and public key.  
    - The below assumes the keys are located in your home SSH directory (~/.ssh) and called `mykey`(private key) and `mykey.pub` (public key)

#### Terraform Deploy Process
1. Create a new directory 
1. With your preferred text editor, copy the below [terraform code](#terraform-code) into a new file and save as `main.tf`.
1. From a terminal run the below commands:
    1. `terraform init`
    1. `terraform apply -var="admin_ssh_key_pub=$(cat ~/.ssh/mykey.pub)" -var="aws_profile='default'`  (Replace the public key path and aws_profile name as needed)
    1. When prompted, enter `yes` to create all the defined resources.
1. Terraform will output the public IP addresses of the two EC2 instances.  Connect to both instances with a command like `ssh -i ~/.ssh/mykey ubuntu@ipaddress` and proceed with the [inital agent setup]({{<relref "/getting-started/trial/base-setup#step-1---install-first-agent">}})

{{<alert color="warning">}}Be sure to [Terraform cleanup process](#terraform-cleanup) below to remove all resources to avoid incurring AWS costs{{</alert>}}

#### Terraform Code
<pre style="max-height: 400px; max-width: 90%" class="line-numbers language-hcl">
<code>terraform {
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = ">= 5.31.0"
        }
    }
}


provider "aws" {
  alias  = "agent1"
  region = var.aws_region_agent1
  profile = var.aws_profile
}

provider "aws" {
  alias  = "agent2"
  region = var.aws_region_agent2
  profile = var.aws_profile
}


variable "aws_profile" {
  type = string
  default = "default"
  description = "Name of configured AWS Profile using the aws cli command"
}

variable "aws_region_agent1" {
    type = string
    default = "us-east-1"
    description = "AWS region to create first agent"
}

variable "aws_region_agent2" {
    type = string
    default = "us-west-1"
    description = "AWS region to create second agent"
}

variable "agent1_name_prefix" {
    type = string
    default = "agent1"
    description = "Name useds for creating all aws resources related to agent1"  
}

variable "agent2_name_prefix" {
    type = string
    default = "agent2"
    description = "Name useds for creating all aws resources related to agent2"  
}

variable "admin_ssh_key_pub" {
    type = string
    description = "Public key added to EC2 instances authorized_keys for ssh access"
  
}

variable "instance_type" {
    type = string
    default = "t2.micro"
    description = "EC2 Instance Type for the deployed instances to host the Trustgrid agent"
  
}


resource "aws_key_pair" "agent1_key_pair" {
    provider = aws.agent1
    key_name   = "${var.agent1_name_prefix}-key"
    public_key = var.admin_ssh_key_pub
}

resource "aws_vpc" "agent1_vpc" {
    provider = aws.agent1
    cidr_block = "10.0.0.0/16"
    enable_dns_support = "true" #gives you an internal domain name
    enable_dns_hostnames = "true" #gives you an internal host name
    instance_tenancy = "default"
    tags = {
        Name = "${var.agent1_name_prefix}-vpc"
    }
}

resource "aws_subnet" "agent1_subnet_public_1" {
    provider = aws.agent1
    vpc_id = aws_vpc.agent1_vpc.id
    cidr_block = "10.0.1.0/24"
    map_public_ip_on_launch = "true" //it makes this a public subnet
    tags = {
        Name = "${var.agent1_name_prefix}-subnet-public-1"
    }
}

resource "aws_route_table" "agent1_subnet_public_1_rt" {
    provider = aws.agent1
    vpc_id = aws_vpc.agent1_vpc.id
    tags = {
      Name = "${var.agent1_name_prefix}-subnet-public-1-rt"
    }  
}

resource "aws_route_table_association" "agent1_subnet_public_1_rt_assoc" {
    provider = aws.agent1
    subnet_id = aws_subnet.agent1_subnet_public_1.id
    route_table_id = aws_route_table.agent1_subnet_public_1_rt.id
  
}

resource "aws_internet_gateway" "agent1_igw" {
    provider = aws.agent1
    vpc_id = aws_vpc.agent1_vpc.id
    tags = {
        Name = "${var.agent1_name_prefix}-igw"
    }
  
}

resource "aws_route" "agent1_internet" {
    provider = aws.agent1
    route_table_id = aws_route_table.agent1_subnet_public_1_rt.id
    destination_cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.agent1_igw.id
  
}

resource "aws_default_security_group" "agent1_default" {
    provider = aws.agent1
    vpc_id = aws_vpc.agent1_vpc.id
    tags = {
        Name = "${var.agent1_name_prefix}-default-sg"
    }
}

resource "aws_vpc_security_group_egress_rule" "agent1_allow_all_out" {
    provider = aws.agent1
    cidr_ipv4 = "0.0.0.0/0"
    from_port = "-1"
    to_port = "-1"
    ip_protocol = "-1"
    security_group_id = aws_default_security_group.agent1_default.id
    description = "Allow all traffic outbound"
}

resource "aws_vpc_security_group_ingress_rule" "agent1_allow_all_vpc" {
    provider = aws.agent1
    ip_protocol = "-1"
    from_port = "-1"
    to_port = "-1"
    referenced_security_group_id = aws_default_security_group.agent1_default.id
    security_group_id = aws_default_security_group.agent1_default.id
    description = "Allow all traffic within VPC"
}

resource "aws_vpc_security_group_ingress_rule" "agent1_allow_ssh" {
    provider = aws.agent1
    ip_protocol = "tcp"
    from_port = 22
    to_port = 22
    cidr_ipv4 = "0.0.0.0/0"
    description = "Allow SSH from any IP. NOT RECOMMENDED FOR PRODUCTION"
    security_group_id = aws_default_security_group.agent1_default.id  
}


data "aws_ami" "agent1_ubuntu" {
    provider = aws.agent1
    most_recent = true

    filter {
      name   = "name"
      values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
    }

    filter {
      name   = "virtualization-type"
      values = ["hvm"]
    }

    owners = ["099720109477"] # Canonical's owner ID
}


resource "aws_instance" "agent1" {
    provider = aws.agent1
    ami = data.aws_ami.agent1_ubuntu.id
    instance_type = var.instance_type
    associate_public_ip_address = true
    subnet_id = aws_subnet.agent1_subnet_public_1.id
    key_name = aws_key_pair.agent1_key_pair.key_name
    security_groups = [ aws_default_security_group.agent1_default.id ]
    tags = {
        Name = var.agent1_name_prefix
    }
    lifecycle {
      ignore_changes = [ security_groups, tags ]
    }
}

output "agent1_public_ip" {
    value = aws_instance.agent1.public_ip
    description = "Public IP address for the agent1 EC2 instances"
}

resource "aws_key_pair" "agent2_key_pair" {
    provider = aws.agent2
    key_name   = "${var.agent2_name_prefix}-key"
    public_key = var.admin_ssh_key_pub
}

resource "aws_vpc" "agent2_vpc" {
    provider = aws.agent2
    cidr_block = "10.0.0.0/16"
    enable_dns_support = "true" #gives you an internal domain name
    enable_dns_hostnames = "true" #gives you an internal host name
    instance_tenancy = "default"
    tags = {
        Name = "${var.agent2_name_prefix}-vpc"
    }
}

resource "aws_subnet" "agent2_subnet_public_1" {
    provider = aws.agent2
    vpc_id = aws_vpc.agent2_vpc.id
    cidr_block = "10.0.1.0/24"
    map_public_ip_on_launch = "true" //it makes this a public subnet
    tags = {
        Name = "${var.agent2_name_prefix}-subnet-public-1"
    }
}

resource "aws_route_table" "agent2_subnet_public_1_rt" {
    provider = aws.agent2
    vpc_id = aws_vpc.agent2_vpc.id
    tags = {
      Name = "${var.agent2_name_prefix}-subnet-public-1-rt"
    }  
}

resource "aws_route_table_association" "agent2_subnet_public_1_rt_assoc" {
    provider = aws.agent2
    subnet_id = aws_subnet.agent2_subnet_public_1.id
    route_table_id = aws_route_table.agent2_subnet_public_1_rt.id
  
}

resource "aws_internet_gateway" "agent2_igw" {
    provider = aws.agent2
    vpc_id = aws_vpc.agent2_vpc.id
    tags = {
        Name = "${var.agent2_name_prefix}-igw"
    }
  
}

resource "aws_route" "agent2_internet" {
    provider = aws.agent2
    route_table_id = aws_route_table.agent2_subnet_public_1_rt.id
    destination_cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.agent2_igw.id
  
}

resource "aws_default_security_group" "agent2_default" {
    provider = aws.agent2
    vpc_id = aws_vpc.agent2_vpc.id
    tags = {
        Name = "${var.agent2_name_prefix}-default-sg"
    }
}

resource "aws_vpc_security_group_egress_rule" "agent2_allow_all_out" {
    provider = aws.agent2
    cidr_ipv4 = "0.0.0.0/0"
    from_port = "-1"
    to_port = "-1"
    ip_protocol = "-1"
    security_group_id = aws_default_security_group.agent2_default.id
    description = "Allow all traffic outbound"
}

resource "aws_vpc_security_group_ingress_rule" "agent2_allow_all_vpc" {
    provider = aws.agent2
    ip_protocol = "-1"
    from_port = "-1"
    to_port = "-1"
    referenced_security_group_id = aws_default_security_group.agent2_default.id
    security_group_id = aws_default_security_group.agent2_default.id
    description = "Allow all traffic within VPC"
}

resource "aws_vpc_security_group_ingress_rule" "agent2_allow_ssh" {
    provider = aws.agent2
    ip_protocol = "tcp"
    from_port = 22
    to_port = 22
    cidr_ipv4 = "0.0.0.0/0"
    description = "Allow SSH from any IP. NOT RECOMMENDED FOR PRODUCTION"
    security_group_id = aws_default_security_group.agent2_default.id  
}


data "aws_ami" "agent2_ubuntu" {
    provider = aws.agent2
    most_recent = true

    filter {
      name   = "name"
      values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
    }

    filter {
      name   = "virtualization-type"
      values = ["hvm"]
    }

    owners = ["099720109477"] # Canonical's owner ID
}


resource "aws_instance" "agent2" {
    provider = aws.agent2
    ami = data.aws_ami.agent2_ubuntu.id
    instance_type = var.instance_type
    associate_public_ip_address = true
    subnet_id = aws_subnet.agent2_subnet_public_1.id
    key_name = aws_key_pair.agent2_key_pair.key_name
    security_groups = [ aws_default_security_group.agent2_default.id ]
    tags = {
        Name = var.agent2_name_prefix
    }
    lifecycle {
      ignore_changes = [ security_groups, tags ]
    }
}

output "agent2_public_ip" {
    value = aws_instance.agent2.public_ip
    description = "Public IP address for the agent2 EC2 instances"
}
</code>
</pre>
 
 {{<alert color="info">}} 
 The terraform code above includes a number of variables with default values that can be overridden if desired. Most notably `aws_region_agent1` and `aws_region_agent2` which control where each of the agents is deployed.
 Feel free to adjust these variables as desired.
 {{</alert>}}

 #### Terraform Cleanup

 1. Open a terminal and navigate to the directory you created the `main.tf` Terraform file during the initial deployment.
 1. Run the command `terraform destroy -var="admin_ssh_key_pub=$(cat ~/.ssh/mykey.pub)" -var="aws_profile='default'`.
 1. When prompted enter `yes`.

 Terraform will proceed to terminate the EC2 instances and delete all the associated VPC resources.

