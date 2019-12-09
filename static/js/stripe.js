var stripe = Stripe('pk_live_7SJOzG0qoGnIjMHbEOnnhEce');
var oneTimeSKU = 'sku_FZzomJKEpaGTVk';
var recurringSKU = 'plan_FZzqRxfQtAx34O';
// var stripe = Stripe('pk_test_U7UN5bE4K6xwGA1nL6kjBfib');
// var oneTimeSKU = 'sku_FZzIBb9nkPEwrd';
// var recurringSKU = 'plan_FZzalXk9lFd7sm';

document.getElementById('donate-once-button').addEventListener('click', function () {
    var amount = +(document.getElementById('donation-amount').value);
    stripe.redirectToCheckout({
        items: [{ sku: oneTimeSKU, quantity: amount }],
        successUrl: 'https://syncthing.net/donations/success/',
        cancelUrl: 'https://syncthing.net/donations/cancelled/',
        submitType: 'donate',
    }).then(function (result) {
        if (result.error) {
            var displayError = document.getElementById('error-message');
            displayError.textContent = result.error.message;
        }
    });
});

document.getElementById('donate-monthly-button').addEventListener('click', function () {
    var amount = +(document.getElementById('donation-amount').value);
    stripe.redirectToCheckout({
        items: [{ plan: recurringSKU, quantity: amount }],
        successUrl: 'https://syncthing.net/donations/recurring/',
        cancelUrl: 'https://syncthing.net/donations/cancelled/',
    }).then(function (result) {
        if (result.error) {
            var displayError = document.getElementById('error-message');
            displayError.textContent = result.error.message;
        }
    });
});