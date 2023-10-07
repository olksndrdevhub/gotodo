
function createNotification(notification) {
    // Clone the template
    const element = htmx.find("[data-notification-template]").cloneNode(true)
    const success_icon_svg = '<svg class="w-6 h-6" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20"><path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 8.207-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414L9 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414Z"/></svg>'
    const warning_icon_svg = '<svg class="w-6 h-6" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20"><path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM10 15a1 1 0 1 1 0-2 1 1 0 0 1 0 2Zm1-4a1 1 0 0 1-2 0V6a1 1 0 0 1 2 0v5Z"/></svg>'
    const danger_icon_svg = '<svg class="w-6 h-6" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20"><path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 11.793a1 1 0 1 1-1.414 1.414L10 11.414l-2.293 2.293a1 1 0 0 1-1.414-1.414L8.586 10 6.293 7.707a1 1 0 0 1 1.414-1.414L10 8.586l2.293-2.293a1 1 0 0 1 1.414 1.414L11.414 10l2.293 2.293Z"/></svg>'
    const succes_icon_classes = ' text-action-success bg-green-100 rounded-lg'
    const warning_icon_classes = ' text-action-warning bg-yellow-100 rounded-lg'
    const danger_icon_classes = ' text-action-danger bg-red-100'
    
    // Remove the data-notification-template attribute
    delete element.dataset.notificationTemplate
  
    // Set the CSS class
    // element.id += notification.tags
    element.classList.remove('hidden');
    element.classList.add('flex');
  
    // Set the text
    htmx.find(element, "[data-notification-body]").innerText = notification.message
    if (notification.type === 'success') {
      element.classList += ' border-action-success'
      // element.style.borderColor = 'rgb(var(--color-action-success) / 1)'
      htmx.find(element, "[data-icon-container]").innerHTML = success_icon_svg
      htmx.find(element, "[data-icon-container]").classList += succes_icon_classes
    } else if (notification.type === 'warning') {
      element.classList += ' border-action-warning'
      htmx.find(element, "[data-icon-container]").innerHTML = warning_icon_svg
      htmx.find(element, "[data-icon-container]").classList += warning_icon_classes
    } else if (notification.type === 'danger') {
      element.classList += ' border-action-danger'
      htmx.find(element, "[data-icon-container]").innerHTML = danger_icon_svg
      htmx.find(element, "[data-icon-container]").classList += danger_icon_classes
    }
    // Add the new element to the container
    htmx.find("[data-notification-container]").appendChild(element)  
  }
  
  function _removeFadeOut( el, speed ) {
    var seconds = speed/1000;
    el.style.transition = "opacity "+seconds+"s ease";
    el.style.opacity = 0;
    setTimeout(function() {
      if (document.body.contains(el)) {
        el.parentNode.removeChild(el);
      }
    }, speed);
  }
  
  function removeNotification(target) {
    let notificationElement = target.closest('.notification');
    _removeFadeOut(notificationElement, 1000);
  }
  
  function removeAllNotifications(){
    htmx.findAll(".notification:not([data-notification-template])").forEach((element) => {
      setTimeout(function() {
        // _removeFadeOut(element, 1000)
      }, 5000)
    })
  }