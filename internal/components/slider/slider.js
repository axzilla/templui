(function () {
  // IIFE
  function initSlider(sliderInput) {
    if (sliderInput.hasAttribute("data-initialized")) return;

    sliderInput.setAttribute("data-initialized", "true");

    const sliderId = sliderInput.id;
    if (!sliderId) return;

    const valueElements = document.querySelectorAll(
      `[data-slider-value][data-slider-value-for="${sliderId}"]`
    );

    function updateValues() {
      valueElements.forEach((el) => {
        el.textContent = sliderInput.value;
      });
    }

    updateValues();
    sliderInput.addEventListener("input", updateValues);
  }

  function initAllComponents(root = document) {
    if (
      root instanceof Element &&
      root.matches('input[type="range"][data-slider-input]')
    ) {
      initSlider(root);
    }
    for (const slider of root.querySelectorAll(
      'input[type="range"][data-slider-input]:not([data-initialized])'
    )) {
      initSlider(slider);
    }
  }

  const handleHtmxSwap = (event) => {
    let target;
    if (event.type === "htmx:afterSwap") {
      target = event.detail.elt;
    }
    if (event.type === "htmx:oobAfterSwap") {
      target = event.detail.target;
    }
    if (target instanceof Element) {
      requestAnimationFrame(() => initAllComponents(target));
    }
  };

  document.addEventListener("DOMContentLoaded", () => initAllComponents());
  document.body.addEventListener("htmx:afterSwap", handleHtmxSwap);
  document.body.addEventListener("htmx:oobAfterSwap", handleHtmxSwap);
})(); // End of IIFE
