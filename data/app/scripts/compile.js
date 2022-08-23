function shareResult() {
    // window.Telegram.WebApp.sendData(localStorage.getItem("result"));
    console.log("send data:");
    console.log(window.Telegram.WebApp.sendData("1"));
}

function setupCompileButtons() {
    window.Telegram.WebApp.MainButton.setText("SHARE");
    window.Telegram.WebApp.MainButton.show();
    window.Telegram.WebApp.MainButton.onClick(shareResult);

    window.Telegram.WebApp.BackButton.onClick(backToCanvas);
    window.Telegram.WebApp.BackButton.show();
}

function backToCanvas() {
    window.Telegram.WebApp.MainButton.offClick(shareResult);
    window.Telegram.WebApp.MainButton.hide();

    window.Telegram.WebApp.BackButton.offClick(backToCanvas);
    window.Telegram.WebApp.BackButton.hide();

    initCanvas();
    Reveal.up();
}

function initCompile() {
    setupCompileButtons();

    const image = document.querySelector("#result");
    image.src = localStorage.getItem("result");
}