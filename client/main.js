const stripe = Stripe($("#publicKey").text());

$("#amount").keyup(function () {
    if ($("#amount").val() > 0) {
        $("#submit").prop("disabled", false);
    } else {
        $("#submit").prop("disabled", true);
    }
})

$("#submit").click(function () {
    var settings = {
        "url": $("#recur").val() ? "/recur" : "/donate",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        "data": {
            "amount": $("#amount").val() * 100
        }
    };

    $.ajax(settings).done(function (response) {
        stripe.redirectToCheckout({
            sessionId: response.sessionID
        }).then(function (result) {
            alert(result.error.message)
        });
    });
})