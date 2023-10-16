
export function initSubscribe(form: HTMLFormElement) {

    const button = form.querySelector("button[type=submit]") as HTMLButtonElement;
    const input = form.querySelector("input[type=email]") as HTMLButtonElement;

    if (!button) {
        throw new Error("form without button");
    }

    button.toggleAttribute("disabled", false);
    input.focus();

    form.addEventListener("htmx:beforeRequest", function() {
        button.toggleAttribute("disabled", true);
    })

    form.addEventListener("htmx:beforeSwap", function() {
        button.toggleAttribute("disabled", true);
    })
}

