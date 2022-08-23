var onGalleryNodes = false;
const images = (()=>{
    return Array.from({length: 10}, (_, i)=>{
        let s = "" + (i+1);
        s = (s.length == 1) ? "0" + s : s;

        return {
            src: "images/" + s + ".jpg",
            id: "el-" + s
        };
    });
})();

function setImg(imgSrc) {
    if (!window.Telegram.WebApp.MainButton.isVisible) {
        window.Telegram.WebApp.MainButton.show();
    }
    localStorage.setItem("image", imgSrc)
}

function goToCanvas() {
    window.Telegram.WebApp.MainButton.hide();
    window.Telegram.WebApp.MainButton.offClick(goToCanvas);

    setTimeout(()=>{
        initCanvas();
        Reveal.down();
    },1)
}

function setupGalleryButtons() {
    window.Telegram.WebApp.MainButton.setText("SIGN");
    window.Telegram.WebApp.MainButton.onClick(goToCanvas)
    window.Telegram.WebApp.MainButton.hide();
}

function setupGalleryNodes() {
    const parent = document.querySelector(".app-gallery");
    for (let img of images) {
        const card = document.createElement("div");
        card.className = "app-gallery-card";
        card.innerHTML = `
            <input 
                type="radio" 
                name="gallery" 
                id="${img.id}"
                value="${img.src}"
            >
            <label 
                for="${img.id}"
                style="background-image: url(${img.src})"
                onclick="setImg('${img.src}')"
            ></label>
        `;

        parent.appendChild(card);
    }

    onGalleryNodes = true;
}

function initGallery() {
    setupGalleryButtons();
    if (!onGalleryNodes) setupGalleryNodes();
}