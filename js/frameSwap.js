function updateFrame(fps) {
    // Get the elements
    var view = document.getElementById("view");
    var nextBufferFrame = document.getElementById("next-buffer-frame");
    var buffer = document.getElementById("buffer");
//    var nextBufferFrame = buffer.querySelector(".buffer-frame");

    // Check if the elements exist
    if (view && nextBufferFrame && buffer) {
        // Move the next buffer frame to the view
        view.appendChild(nextBufferFrame);

        requestAnimationFrame(() => {
            nextBufferFrame.classList.remove("hidden");
        });

        var oldFrame = view.querySelector(".frame:not(.hidden)");
        if (oldFrame) {
            oldFrame.classList.add("hidden");
            setTimeout(() => {
                if (oldFrame.parentNode) {
                    oldFrame.parentNode.removeChild(oldFrame);
                }
            }, (1000 / fps));
        }

        nextBufferFrame.removeAttribute("id");

        // Set the ID on the next buffer frame
        var bufferFrame = buffer.querySelector(".buffer-frame:not([id])");
        if (bufferFrame) {
            bufferFrame.setAttribute("id", "next-buffer-frame");
        }
    } else {
        console.error("One or more elements not found:", { view, nextBufferFrame, buffer });
    }

    // Schedule the next update
    setTimeout(updateFrame, 1000 / fps, fps);
}

// Start the update loop
//setTimeout(updateFrame, 10);

