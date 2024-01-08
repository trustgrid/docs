function toggleFullScreen(imgElement) {
    if (!document.fullscreenElement) {
      if (imgElement.requestFullscreen) {
        imgElement.requestFullscreen();
      } else if (imgElement.webkitRequestFullscreen) { // Safari
        imgElement.webkitRequestFullscreen();
      } else if (imgElement.msRequestFullscreen) { // IE11
        imgElement.msRequestFullscreen();
      }
    } else {
      if (document.exitFullscreen) {
        document.exitFullscreen();
      } else if (document.webkitExitFullscreen) { // Safari
        document.webkitExitFullscreen();
      } else if (document.msExitFullscreen) { // IE11
        document.msExitFullscreen();
      }
    }
  }
  
  function toggleImageSize(imgElement) {
    if (imgElement.style.maxWidth === '90%' && imgElement.style.maxHeight === '90%') {
        imgElement.style.position = '';
        imgElement.style.maxWidth = '';
        imgElement.style.maxHeight = '';
        imgElement.style.top = '';
        imgElement.style.left = '';
        imgElement.style.transform = '';
        imgElement.style.zIndex = '';
        imgElement.style.boxShadow = ''; // Remove shadow
    } else {
        imgElement.style.position = 'fixed';
        imgElement.style.maxWidth = '90%';
        imgElement.style.maxHeight = '90%';
        imgElement.style.top = '50%';
        imgElement.style.left = '50%';
        imgElement.style.transform = 'translate(-50%, -50%)';
        imgElement.style.zIndex = '1000';
        imgElement.style.boxShadow = '0px 0px 20px 5px rgba(0, 0, 0, 0.5)'; // Add shadow
    }
}

  