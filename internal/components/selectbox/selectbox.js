(function () {
  let isSelecting = false;

  function checkVisibility(element) {
    const style = window.getComputedStyle(element);
    return (
      style.display !== "none" &&
      style.visibility !== "hidden" &&
      style.opacity !== "0"
    )
  };

  function initSelect(wrapper) {
    if (!wrapper || wrapper.hasAttribute("data-initialized")) return;
    wrapper.setAttribute("data-initialized", "true");

    const triggerButton = wrapper.querySelector("button.select-trigger");
    if (!triggerButton) {
      console.error(
        "Select box: Trigger button (.select-trigger) not found in wrapper",
        wrapper,
      );
      return;
    }

    const contentID = triggerButton.getAttribute(
      "data-tui-selectbox-content-id",
    );
    const content = contentID ? document.getElementById(contentID) : null;
    const valueEl = triggerButton.querySelector(".select-value");
    const hiddenInput = triggerButton.querySelector('input[type="hidden"]');
    if (!content || !valueEl || !hiddenInput) {
      console.error(
        "Select box: Missing required elements for initialization.",
        {
          wrapper,
          contentID,
          contentExists: !!content,
          valueElExists: !!valueEl,
          hiddenInputExists: !!hiddenInput,
        },
      );
      return;
    }

    const isMultiple =
      triggerButton.getAttribute("data-tui-selectbox-multiple") === "true";
    const showPills =
      triggerButton.getAttribute("data-tui-selectbox-show-pills") === "true";
    const searchInput = content.querySelector("[data-tui-selectbox-search]");
    const form = wrapper.closest("form");

    // Remove existing event listeners
    document.removeEventListener("click", handleTriggerClickFocusSearch);
    triggerButton.removeEventListener("keydown", handleTriggerKeydownOpenContent);
    triggerButton.removeEventListener("keydown", handleTriggerKeydownFocusSearch);
    content.removeEventListener("keydown", handleContentKeydownNavigation);
    content.removeEventListener("click", handleContentClickSelect);
    content.removeEventListener("keydown", handleContentKeydownSelect);
    content.removeEventListener("mouseover", handleContentMouseoverStyle);
    content.removeEventListener("mouseleave", handleContentMouseleaveResetStyle);
    if (searchInput) {
      searchInput.removeEventListener("input", handleSearchInputOnInput);
    }
    if (form) {
      form.removeEventListener("reset", handleFormReset);
    }

    // EventListener Functions
    function handleTriggerKeydownOpenContent(event) {
      if (
        event.key.length === 1 ||
        event.key === "Backspace" ||
        event.key === "Delete"
      ) {
        event.preventDefault();
        document.getElementById(contentID).click();
        setTimeout(() => {
          if (searchInput) {
            searchInput.focus();
            if (event.key !== "Backspace" && event.key !== "Delete") {
              searchInput.value = event.key;
            }
          }
        }, 0);
      }
    }

    function handleTriggerClickFocusSearch(event) {
      if (triggerButton.contains(event.target)) {
        setTimeout(() => checkVisibility(content) && searchInput.focus(), 50);
      }
    }

    function handleTriggerKeydownFocusSearch(event) {
      if (event.key === "Enter" || event.key === " ") {
        setTimeout(() => checkVisibility(content) && searchInput.focus(), 50);
      }
    }

    function handleSearchInputOnInput(event) {
      const searchTerm = event.target.value.toLowerCase().trim();
      const items = content.querySelectorAll(".select-item");

      items.forEach((item) => {
        const itemText =
          item
            .querySelector(".select-item-text")
            ?.textContent.toLowerCase() || "";
        const itemValue =
          item.getAttribute("data-tui-selectbox-value")?.toLowerCase() || "";
        const isVisible =
          searchTerm === "" ||
          itemText.includes(searchTerm) ||
          itemValue.includes(searchTerm);

        item.style.display = isVisible ? "" : "none";
      });
    }

    function handleContentKeydownNavigation(event) {
      if (event.key === "ArrowDown" || event.key === "ArrowUp") {
        event.preventDefault();
        const visibleItems = Array.from(
          content.querySelectorAll(".select-item"),
        ).filter((item) => item.style.display !== "none");

        if (visibleItems.length === 0) return;

        const currentFocused = content.querySelector(".select-item:focus");
        let nextIndex = 0;

        if (currentFocused) {
          const currentIndex = visibleItems.indexOf(currentFocused);
          if (event.key === "ArrowDown") {
            nextIndex = (currentIndex + 1) % visibleItems.length;
          } else {
            nextIndex =
              (currentIndex - 1 + visibleItems.length) % visibleItems.length;
          }
        }

        visibleItems[nextIndex].focus();
      } else if (event.key === "Enter") {
        event.preventDefault();
        const focusedItem = content.querySelector(".select-item:focus");
        if (focusedItem) {
          selectItem(focusedItem);
        }
      } else if (event.key === "Escape") {
        event.preventDefault();
        const focusedItem = content.querySelector(".select-item:focus");
        if (focusedItem) {
          // If focus is on an item, move to search input
          searchInput.focus();
        } else if (document.activeElement === searchInput) {
          // If focus is on search input, close popover and return to trigger
          if (window.closePopover) {
            window.closePopover(contentID, true);
            setTimeout(() => {
              triggerButton.focus();
            }, 50);
          }
        }
      }
    }

    function handleContentClickSelect(event) {
      const item = event.target.closest(".select-item");
      if (item) selectItem(item);
    }

    function handleContentKeydownSelect(event) {
      const item = event.target.closest(".select-item");
      if (item && (event.key === "Enter" || event.key === " ")) {
        event.preventDefault();
        selectItem(item);
      }
    }

    function handleContentMouseoverStyle(event) {
      const item = event.target.closest(".select-item");
      if (!item || item.getAttribute("data-tui-selectbox-disabled") === "true")
        return;
      // Reset all others first
      content.querySelectorAll(".select-item").forEach((el) => {
        el.classList.remove("bg-accent", "text-accent-foreground", "bg-muted");
      });
      // Apply hover style only if not selected
      if (item.getAttribute("data-tui-selectbox-selected") !== "true") {
        item.classList.add("bg-accent", "text-accent-foreground");
      }
    }

    // Reset visual state of items
    function handleContentMouseleaveResetStyle() {
      content.querySelectorAll(".select-item").forEach((item) => {
        if (item.getAttribute("data-tui-selectbox-selected") === "true") {
          item.classList.add("bg-accent", "text-accent-foreground");
          item.classList.remove("bg-muted");
        } else {
          item.classList.remove(
            "bg-accent",
            "text-accent-foreground",
            "bg-muted",
          );
        }
      });
    }

    function handleFormReset(){
      // Clear all selections
      content.querySelectorAll(".select-item").forEach((item) => {
        item.setAttribute("data-tui-selectbox-selected", "false");
        const check = item.querySelector(".select-check");
        if (check) check.classList.replace("opacity-100", "opacity-0");
        item.classList.remove("bg-accent", "text-accent-foreground");
      });

      // Reset display value to placeholder
      const placeholder =
        valueEl.getAttribute("data-placeholder") ||
        triggerButton
          .querySelector(".select-value")
          ?.getAttribute("data-placeholder") ||
        "Select...";
      valueEl.textContent = placeholder;
      valueEl.classList.add("text-muted-foreground");

      // Clear hidden input
      hiddenInput.value = "";

      // Clear search input if exists
      if (searchInput) {
        searchInput.value = "";
        // Show all items again
        content.querySelectorAll(".select-item").forEach((item) => {
          item.style.display = "";
        });
      }
    }

    // Add keyboard event handler for trigger button
    triggerButton.addEventListener("keydown", handleTriggerKeydownOpenContent);

    if (searchInput) {
      // Focus when opened by click
      document.addEventListener("click", handleTriggerClickFocusSearch);
      // Focus when opened by Enter key
      triggerButton.addEventListener("keydown", handleTriggerKeydownFocusSearch);
      searchInput.addEventListener("input", handleSearchInputOnInput);
    }

    // Keyboard navigation event listener
    content.addEventListener("keydown", handleContentKeydownNavigation);

    // Initialize display value if an item is pre-selected
    const selectedItems = content.querySelectorAll(
      '.select-item[data-tui-selectbox-selected="true"]',
    );
    if (selectedItems.length > 0) {
      if (isMultiple) {
        if (showPills) {
          valueEl.innerHTML = "";
          const pillsContainer = document.createElement("div");
          pillsContainer.className =
            "flex flex-nowrap overflow-hidden max-w-full whitespace-nowrap gap-1";
          Array.from(selectedItems).forEach((selectedItem) => {
            const pill = document.createElement("div");
            pill.className =
              "flex items-center gap-1 px-2 py-1 text-xs rounded-md bg-primary text-primary-foreground";
            const pillText = document.createElement("span");
            pillText.textContent =
              selectedItem.querySelector(".select-item-text").textContent;
            const closeButton = document.createElement("button");
            closeButton.className = "hover:text-destructive focus:outline-none";
            closeButton.innerHTML = "x";
            closeButton.onclick = (e) => {
              e.stopPropagation();
              selectItem(selectedItem);
            };
            pill.appendChild(pillText);
            pill.appendChild(closeButton);
            pillsContainer.appendChild(pill);
          });
          valueEl.appendChild(pillsContainer);
          valueEl.classList.remove("text-muted-foreground");

          // Pills overflow control
          setTimeout(() => {
            const pillsWidth = pillsContainer.scrollWidth;
            const valueWidth = valueEl.clientWidth;
            if (pillsWidth > valueWidth) {
              const selectedCountText =
                triggerButton.getAttribute(
                  "data-tui-selectbox-selected-count-text",
                ) || `${selectedItems.length} items selected`;
              const msg = selectedCountText.replace(
                "{n}",
                selectedItems.length,
              );
              valueEl.innerHTML = msg;
              valueEl.classList.remove("text-muted-foreground");
            }
          }, 0);
        } else {
          valueEl.textContent = `${selectedItems.length} items selected`;
          valueEl.classList.remove("text-muted-foreground");
        }
        // Store selected values as CSV
        const selectedValues = Array.from(selectedItems).map((item) =>
          item.getAttribute("data-tui-selectbox-value"),
        );
        hiddenInput.value = selectedValues.join(",");
      } else {
        // For single selection, show the selected item's text
        const selectedItem = selectedItems[0];
        const itemText = selectedItem.querySelector(".select-item-text");
        if (itemText) {
          valueEl.textContent = itemText.textContent;
          valueEl.classList.remove("text-muted-foreground");
        }
        if (hiddenInput) {
          const value = selectedItem.getAttribute("data-tui-selectbox-value") || "";
          // Only set initial value if not already set
          if (!hiddenInput.hasAttribute("data-initialized")) {
            hiddenInput.value = value;
            hiddenInput.setAttribute("data-initialized", "true");
            hiddenInput.dispatchEvent(new Event("change", { bubbles: true }));
          }
        }
      }
    }

    // Select an item
    function selectItem(item) {
      if (
        !item ||
        item.getAttribute("data-tui-selectbox-disabled") === "true" ||
        isSelecting
      )
        return;

      isSelecting = true;

      const value = item.getAttribute("data-tui-selectbox-value");
      const itemText = item.querySelector(".select-item-text");

      if (isMultiple) {
        // Toggle selection for multiple mode
        const isSelected =
          item.getAttribute("data-tui-selectbox-selected") === "true";
        item.setAttribute(
          "data-tui-selectbox-selected",
          (!isSelected).toString(),
        );

        if (!isSelected) {
          item.classList.add("bg-accent", "text-accent-foreground");
          const check = item.querySelector(".select-check");
          if (check) check.classList.replace("opacity-0", "opacity-100");
        } else {
          item.classList.remove("bg-accent", "text-accent-foreground");
          const check = item.querySelector(".select-check");
          if (check) check.classList.replace("opacity-100", "opacity-0");
        }

        // Update display value
        const selectedItems = content.querySelectorAll(
          '.select-item[data-tui-selectbox-selected="true"]',
        );
        if (selectedItems.length > 0) {
          if (showPills) {
            // Clear existing content
            valueEl.innerHTML = "";

            // Create pills container
            const pillsContainer = document.createElement("div");
            pillsContainer.className =
              "flex flex-nowrap overflow-hidden max-w-full whitespace-nowrap gap-1";

            // Add pills for each selected item
            Array.from(selectedItems).forEach((selectedItem) => {
              const pill = document.createElement("div");
              pill.className =
                "flex items-center gap-1 px-2 py-0.5 text-xs rounded-full bg-accent text-accent-foreground";

              const pillText = document.createElement("span");
              pillText.textContent =
                selectedItem.querySelector(".select-item-text").textContent;

              const closeButton = document.createElement("button");
              closeButton.className =
                "hover:text-destructive focus:outline-none";
              closeButton.innerHTML = "×";
              closeButton.onclick = (e) => {
                e.stopPropagation();
                selectItem(selectedItem);
              };

              pill.appendChild(pillText);
              pill.appendChild(closeButton);
              pillsContainer.appendChild(pill);
            });

            valueEl.appendChild(pillsContainer);
            valueEl.classList.remove("text-muted-foreground");

            // Pills overflow kontrolü
            setTimeout(() => {
              const pillsWidth = pillsContainer.scrollWidth;
              const valueWidth = valueEl.clientWidth;
              if (pillsWidth > valueWidth) {
                const selectedCountText =
                  triggerButton.getAttribute(
                    "data-tui-selectbox-selected-count-text",
                  ) || `${selectedItems.length} items selected`;
                const msg = selectedCountText.replace(
                  "{n}",
                  selectedItems.length,
                );
                valueEl.innerHTML = msg;
                valueEl.classList.remove("text-muted-foreground");
              }
            }, 0);
          } else {
            valueEl.textContent = `${selectedItems.length} items selected`;
            valueEl.classList.remove("text-muted-foreground");
          }
        } else {
          valueEl.textContent =
            valueEl.getAttribute("data-tui-selectbox-placeholder") || "";
          valueEl.classList.add("text-muted-foreground");
        }

        // Update hidden input with CSV of selected values
        const selectedValues = Array.from(selectedItems).map((item) =>
          item.getAttribute("data-tui-selectbox-value"),
        );
        hiddenInput.value = selectedValues.join(",");
        hiddenInput.dispatchEvent(new Event("change", { bubbles: true }));
      } else {
        // Single selection mode
        // Reset all items in this content
        content.querySelectorAll(".select-item").forEach((el) => {
          el.setAttribute("data-tui-selectbox-selected", "false");
          el.classList.remove(
            "bg-accent",
            "text-accent-foreground",
            "bg-muted",
          );
          const check = el.querySelector(".select-check");
          if (check) check.classList.replace("opacity-100", "opacity-0");
        });

        // Mark new selection
        item.setAttribute("data-tui-selectbox-selected", "true");
        item.classList.add("bg-accent", "text-accent-foreground");
        const check = item.querySelector(".select-check");
        if (check) check.classList.replace("opacity-0", "opacity-100");

        // Update display value
        if (valueEl && itemText) {
          valueEl.textContent = itemText.textContent;
          valueEl.classList.remove("text-muted-foreground");
        }

        // Update hidden input & trigger change event
        if (hiddenInput && value !== null) {
          const oldValue = hiddenInput.value;
          hiddenInput.value = value;

          // Only trigger change if value actually changed
          if (oldValue !== value) {
            hiddenInput.dispatchEvent(new Event("change", { bubbles: true }));
          }
        }

        // Close the popover using the correct contentID
        if (window.closePopover) {
          window.closePopover(contentID, true);
          // Return focus to trigger
          setTimeout(() => {
            triggerButton.focus();
          }, 50);
        } else {
          console.warn("closePopover function not found");
        }
      }

      setTimeout(() => {
        isSelecting = false;
      }, 100);
    }


    // Event Listeners for Items (delegated from content for robustness)
    content.addEventListener("click", handleContentClickSelect);
    content.addEventListener("keydown", handleContentKeydownSelect);
    // Event: Mouse hover on items (delegated)
    content.addEventListener("mouseover", handleContentMouseoverStyle);
    // Reset hover styles when mouse leaves the content area
    content.addEventListener("mouseleave", handleContentMouseleaveResetStyle);

    // Form reset support
    if (form) {
      form.addEventListener("reset", handleFormReset);
    }
  }

  function init(root = document) {
    const containers = root.querySelectorAll(".select-container");
    containers.forEach(initSelect);
  }

  window.templUI = window.templUI || {};
  window.templUI.selectbox = { init: init };

  document.addEventListener("DOMContentLoaded", () => init());
})();

