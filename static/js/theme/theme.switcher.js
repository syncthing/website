const head = document.getElementsByTagName("head")[0];
const link = document.getElementById('theme');

linkLoader(getTheme());

function themeSelectorClicked() {
    let next = nextTheme(getTheme());
    linkLoader(next);
    saveTheme(next);
    this.className = getThemeClasses(next);
}

//This function loads a css file according to the value of its parameter
function linkLoader(param) {
    if (param == 'light') {
        link.removeAttribute('href');
        document.querySelector('.logo').setAttribute('src', 'https://www.kastelo.net/img/logo.svg')
    } else if (param == 'dark') {
        link.setAttribute('href', '/css/dark.css');
        document.querySelector('.logo').setAttribute('src', 'https://www.kastelo.net/img/logo-light.svg')
    } else {
        link.setAttribute('href', '/css/auto.css');
    }
    head.appendChild(link);
}

function getTheme() {
    try {
        const theme = localStorage.getItem('theme');
        if (!theme) {
            return 'auto';
        }
        return theme;
    } catch {
        return 'auto';
    }
}

function getThemeClasses(theme) {
    if (theme == 'light') {
        return 'fas fa-sun';
    } else if (theme == 'dark') {
        return 'far fa-moon';
    } else {
        return 'fas fa-magic';
    }
}

function nextTheme(current) {
    if (current == 'light') {
        return 'dark';
    } else if (current == 'dark') {
        return 'auto';
    } else {
        return 'light';
    }
}

function saveTheme(theme) {
    try {
        localStorage.setItem('theme', theme);
    } catch (e) {
        console.log(e);
    }
}
