/**
* @param {HTMLFormElement} form
*/
function disableSubmit(form) {
    /** @type HTMLButtonElement */
    const button = form.querySelector("button[type=submit]");

    if (!button) {
        throw new Error("form without button");
    }

    // ensure at load time that button is not disabled
    button.removeAttribute("disabled");

    form.addEventListener("htmx:beforeRequest", function() {
        button.toggleAttribute("disabled", true);
    })

    form.addEventListener("htmx:beforeSwap", function() {
        button.toggleAttribute("disabled", true);
    })
}

// on document ready
document.addEventListener("DOMContentLoaded", function() {
    console.log("here");
    document.body.addEventListener('htmx:beforeSwap', function(evt) {
        console.log("here 422");
        if (evt.detail.xhr.status === 422) {
            console.log("here 422 YES");
            evt.detail.shouldSwap = true;
            evt.detail.isError = false;
        }
    });
});
