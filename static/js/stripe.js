var stripe = Stripe('pk_live_7SJOzG0qoGnIjMHbEOnnhEce');
//var stripe = Stripe('pk_test_U7UN5bE4K6xwGA1nL6kjBfib');

function redirect(sessionID) {
    stripe.redirectToCheckout({ sessionId: sessionID }).then(function (result) {
        if (result.error) {
            var displayError = document.getElementById('error-message');
            displayError.textContent = result.error.message;
        }
    });
}

document.getElementById('donate-once-button').addEventListener('click', function () {
    var amount = +(document.getElementById('donation-amount').value);
    var req = JSON.stringify({ recurring: false, count: amount });
    $.post("/.netlify/functions/donate", req,
        function (data) {
            redirect(data.sessionID);
        }, "json"
    );
});

document.getElementById('donate-monthly-button').addEventListener('click', function () {
    var amount = +(document.getElementById('donation-amount').value);
    var req = JSON.stringify({ recurring: true, count: amount });
    $.post("/.netlify/functions/donate", req,
        function (data) {
            redirect(data.sessionID);
        }, "json"
    );
});