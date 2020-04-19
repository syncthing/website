var stripe = Stripe('pk_live_7SJOzG0qoGnIjMHbEOnnhEce');
//var stripe = Stripe('pk_test_U7UN5bE4K6xwGA1nL6kjBfib');

var captchaToken = null;
function solvedCaptcha(token) {
    captchaToken = token;
    $('#donate-once-button').removeAttr('disabled');
    $('#donate-monthly-button').removeAttr('disabled');
}

function redirect(sessionID) {
    stripe.redirectToCheckout({ sessionId: sessionID }).then(function (result) {
        if (result.error) {
            $('#error-message').text(result.error.message);
        }
    });
}

document.getElementById('donate-once-button').addEventListener('click', function () {
    $('#donate-once-button').attr('disabled', 'true');
    $('#donate-monthly-button').attr('disabled', 'true');
    var amount = +(document.getElementById('donation-amount').value);
    var req = JSON.stringify({ recurring: false, count: amount, captcha: captchaToken });
    $.post("/.netlify/functions/donate", req,
        function (data) {
            redirect(data.sessionID);
        }, "json"
    ).fail(function () {
        $('#error-message').text("An error ocurred. Please reload the page and try again. :(");
    });
});

document.getElementById('donate-monthly-button').addEventListener('click', function () {
    $('#donate-once-button').attr('disabled', 'true');
    $('#donate-monthly-button').attr('disabled', 'true');
    var amount = +(document.getElementById('donation-amount').value);
    var req = JSON.stringify({ recurring: true, count: amount, captcha: captchaToken });
    $.post("/.netlify/functions/donate", req,
        function (data) {
            redirect(data.sessionID);
        }, "json"
    ).fail(function () {
        $('#error-message').text("An error ocurred. Please reload the page and try again. :(");
    });
});