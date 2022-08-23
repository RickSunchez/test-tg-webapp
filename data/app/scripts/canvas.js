var onCanvasNodes = false;

const w = 500, h = 400;

const bgCanvas = document.querySelector("#bg-canvas");
const drawCanvas = document.querySelector("#draw-canvas");
const sign = document.querySelector("#sign");

const bgctx = bgCanvas.getContext("2d");
const ctx = drawCanvas.getContext("2d");

bgCanvas.width = w;
bgCanvas.height = h;
drawCanvas.width = w;
drawCanvas.height = h;

function getImageSrc() {
    return localStorage.getItem("image");
}

function wrap(context, text, maxWidth) {
    var words = text.split(" ");
    var rows = [];
    var currentRow = words[0];

    for (let i=1; i<words.length; i++) {
        let width = context.measureText(currentRow + " " + words[i]).width;
        if (width < maxWidth) {
            currentRow += " " + words[i];
        } else {
            rows.push(currentRow);
            currentRow = words[i];
        }
    }
    rows.push(currentRow);
    return rows;
}

function compileCanvases() {
    bgctx.drawImage(drawCanvas, 0, 0);
    const dataURL = bgCanvas.toDataURL("image/jpeg");
    localStorage.setItem("result", dataURL);
}

function goToCompile() {
    window.Telegram.WebApp.MainButton.hide();
    window.Telegram.WebApp.MainButton.offClick(goToCompile);

    window.Telegram.WebApp.BackButton.offClick(backToGallery);
    window.Telegram.WebApp.BackButton.hide();

    compileCanvases();
    initCompile();
    Reveal.down();
}

function backToGallery() {
    window.Telegram.WebApp.MainButton.offClick(goToCompile);
    window.Telegram.WebApp.MainButton.hide();

    window.Telegram.WebApp.BackButton.offClick(backToGallery);
    window.Telegram.WebApp.BackButton.hide();

    setTimeout(()=>{
        initGallery();
        Reveal.up();
    },1);
}

function setupCanvasButtons() {
    window.Telegram.WebApp.MainButton.setText("RESULT");
    window.Telegram.WebApp.MainButton.onClick(goToCompile);
    window.Telegram.WebApp.MainButton.hide();

    window.Telegram.WebApp.BackButton.onClick(backToGallery);
    window.Telegram.WebApp.BackButton.show();
}

function setupCanvasImage() {
    const bg = new Image();
    bg.src = getImageSrc();

    bgctx.clearRect(0, 0, w, h);
    bg.onload = () => bgctx.drawImage(bg, 0, 0, w, h);
}

function setupCanvasNodes() {
    sign.addEventListener("input", (ev)=> {
        if (!window.Telegram.WebApp.MainButton.isVisible) {
            window.Telegram.WebApp.MainButton.show();
        }
        ctx.clearRect(0, 0, w, h);

        ctx.fillStyle = "white";
        ctx.strokeStyle = "black";
        ctx.font = "40px Arial Black";

        ctx.textAlign = "center";
        ctx.textBaseline = "center";

        var text = ev.target.value;
        text = text.replace(/\s+/g, " ");
        ev.target.value = text;

        var rows = wrap(ctx, text, 0.8*w);
        var lineHeight = 38;
        var textHeight = lineHeight * rows.length;
        var textStart = h/2 - textHeight/2 + lineHeight;

        for (let i in rows) {
            let x = w/2;
            let y = textStart + lineHeight*i
            ctx.fillText(rows[i], x, y);
            ctx.strokeText(rows[i], x, y);
        }
    })

    onCanvasNodes = true;
}

function initCanvas() {
    setupCanvasButtons();
    setupCanvasImage();
    if (!onCanvasNodes) setupCanvasNodes();
}