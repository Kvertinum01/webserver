welcome_messages = [
    "Привет", "Здравствуйте", "Welcome", "Добро пожаловать",
    "Hello", "¡Hola!", "你好", "안녕하십니까", "Bonjour", "Salut",
]

window.onload = () => {
    var rand = Math.floor(Math.random() * welcome_messages.length);
    var helloElement = document.getElementById("welcome-message");
    helloElement.innerHTML = welcome_messages[rand]

    var header = document.getElementById("header");
    var sticky = header.offsetTop;

    window.onscroll = () => {
        if (window.pageYOffset > sticky) {
            header.classList.add("sticky");
        } else {
            header.classList.remove("sticky");
        }
    }
}

