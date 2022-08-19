<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>WebApp</title>

    <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />
    <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />
    <link rel="preconnect" href="https://fonts.googleapis.com"> 
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin> 
    <link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet">
    
    <script src="https://telegram.org/js/telegram-web-app.js"></script>
    
    <style>
        :root {
            --color-main: #f8a917;
            --color-red: #e64d44;
        }
        body {
            font-family: 'Roboto', sans-serif;
        }
        #order {
            margin: 0px;
            padding: 0px;
            flex-direction: row;
            justify-content: center;
            align-items: center;
            flex-wrap: wrap;
        }
        .material-symbols-rounded {
            font-variation-settings: 'FILL' 1, 'wght' 400, 'GRAD' 0, 'opsz' 48;
            font-size: 4em;
            padding: 4px;
        }
        .material-symbols-outlined {
            font-variation-settings: 'FILL' 0, 'wght' 400, 'GRAD' 0, 'opsz' 48;
            font-size: 2em;
        }

        .app-card {
            position: relative;
            padding: 2px;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            min-width: 80px;
        }
        .app-card[count="0"]:after {
            display: none;
        }
        .app-card:after {
            position: absolute;
            top: 2px;
            right: 2px;
            content: attr(count);
            background-color: var(--color-main);
            display: block;
            width: 22px;
            height: 22px;
            line-height: 22px;
            text-align: center;
            border-radius: 50%;
            color: white;
            font-weight: 900;
        }

        .card-buttons-container {
            display: flex;
            flex-direction: row;
            justify-content: space-around;
            font-weight: 900;
            color: white;
            width: 100%;
        }
        .app-card .card-button {
            background-color: var(--color-main);
            font-size: 1em;
            border-radius: 5px;
            display: flex;
            justify-content: center;
            align-items: center;
            -moz-user-select: none;
            -webkit-user-select: none;
            cursor: pointer;
            width: 38px;
            height: 30px;
        }
            .app-card .card-button[action="remove"] {
                background-color: var(--color-red);
            }
 
        .app-card .card-button[action="add"]:before {
            content: "+";
        }
        .app-card[count="0"] .card-button[action="add"]{
            width: 80px;
        }
        .app-card[count="0"] .card-button[action="add"]:before {
            content: "ADD";
        }
        .app-card[count="0"] .card-button[action="remove"]{
            display: none;
        }

        section {
            display: none;
        }
        section[active] {
            display: flex;
            color: white;
        }
    </style>
</head>
<body>
    <section id="order" active></section>
    <section id="confirm">New Page!</section>
</body>
<script>
    class App {
        constructor(parent){
            this.root = document.querySelector(parent);
            this.cards = [];

            this.root.addEventListener("click", (ev)=>{
                var sum = this.cards.reduce((prev, curr)=>{
                    return prev + curr.count;
                }, 0)

                if (sum > 0) {
                    Telegram.WebApp.MainButton.show()
                } else {
                    Telegram.WebApp.MainButton.hide()
                }
            })
        }

        add(...cards) {
            for (let card of cards) this.cards.push(card);
        }

        nextSection() {
            const active = document.body.querySelector("section[active]");
            active.removeAttribute("active");
            active.nextElementSibling.setAttribute("active", "");
        }
    }

    class Card {
        constructor(parent, type){
            this.parent = document.querySelector(parent);
            this.type = type;
            this._count = 0;

            this.root = document.createElement("div");
            this.root.className = "app-card";
            this.root.setAttribute("count", this._count);

            this.root.innerHTML = this.template();

            this.parent.appendChild(this.root);

            this.root.querySelector(".card-button[action=remove]")
                .addEventListener("click", (ev)=>{
                    this.count -= 1;
                })
            
            this.root.querySelector(".card-button[action=add]")
                .addEventListener("click", (ev)=>{
                    this.count += 1;
                })
        }

        set count(value) {
            this._count = (value < 0) ? 0 : value;
            this.root.setAttribute("count", this._count);
        }
        get count() {
            return this._count;
        }

        template() {
            return `
                <span class="material-symbols-rounded">
                    ${this.type}
                </span>
                <div class="card-buttons-container">
                    <div class="card-button" action="remove">
                        -
                    </div>
                    <div class="card-button" action="add">
                    </div>
                </div>
            `
        }
    }

    const foods = [
        "lunch_dining",
        "cake",
        "emoji_food_beverage",
        "local_pizza",
        "bakery_dining",
        "ramen_dining",
        "icecream",
        "soup_kitchen",
        "kebab_dining",
        "fastfood",
        "local_cafe"
    ]

    const app = new App("#order");
    for (let f of foods)
        app.add(new Card("#order", f))

    Telegram.WebApp.onEvent('mainButtonClicked', function(ev){
        console.log(1);
        console.log(app);
        app.nextSection();
        document.location.replace("next.php");
    });
    Telegram.WebApp.MainButton.setText("ORDER");
</script>
</html>