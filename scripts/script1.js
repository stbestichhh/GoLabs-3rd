// Виконання AJAX запиту для надсилання команди серверу
function sendGreenBorderCommand() {
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/create_green_border", true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
      console.log("Команда на створення зеленої рамки відправлена успішно.");
    }
  };
  xhr.send();
}
