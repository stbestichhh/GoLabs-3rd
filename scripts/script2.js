// Відправлення команди на малювання фігури на сервер
function drawFigure() {
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/draw_figure", true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
      console.log("Фігура намальована успішно.");
      moveFigure();
    }
  };
  xhr.send();
}

// Відправлення команд на переміщення фігури кожну секунду
function moveFigure() {
  setInterval(function () {
    var x = Math.random() * 100; //  генерація координати x
    var y = Math.random() * 100; //  генерація координати y
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/move_figure", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({ x: x, y: y }));
  }, 1000); // Кожну секунду
}
