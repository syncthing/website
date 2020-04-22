---
title: Donations
header: true
---

Developing Syncthing costs money, in domain and hosting fees for the various
servers we need to run the operation such as discovery servers, build
servers and this website. Your donations help fund these costs. We also
periodically award grants for the development of specific features, which is
paid for by these donations.

Donate securely using any major credit or debit card, processed by [Stripe](https://stripe.com).

<div class="g-recaptcha" data-callback="solvedCaptcha" data-sitekey="6LfkXOoUAAAAAJweCDgLlleUUy5dalTbV2mLdruJ"></div>
<form class="form-inline">
    <label class="sr-only" for="donation-amount">Donation Amount</label>
    <div class="input-group mr-sm-2 my-2">
        <div class="input-group-prepend">
            <div class="input-group-text">&euro;</div>
        </div>
        <input type="number" class="form-control text-right input-lg" name="amount" id="donation-amount" placeholder="Amount" min="1" max="1000" value="20" required>
    </div>
    <button type="button" disabled="true" class="btn btn-success mr-sm-2 my-2" id="donate-once-button" role="link"><i class="fa fa-hand-holding-usd"></i>&ensp;Donate Once</button>
    <button type="button" disabled="true" class="btn btn-primary mr-sm-2 my-2" id="donate-monthly-button" role="link"><i class="fa fa-redo-alt"></i>&ensp;Donate Monthly</button>
</form>
<div id="error-message"></div>
<p>
<script src="https://js.stripe.com/v3" defer></script>
<script src="https://www.google.com/recaptcha/api.js" defer></script>
<script src="/js/stripe.js" defer></script>

If you experience any issues with the donation handling, regret your
donation and want a refund, or want to cancel a recurring donation please
feel free to contact [donations@syncthing.org](mailto:donations@syncthing.org) at any time.

If you'd like to become a corporate sponsor of the project and be featured here
we're happy to discuss that too!

{{% sponsors %}}
