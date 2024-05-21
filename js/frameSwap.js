function updateSeg() {
    // Get the elements
    var view = document.getElementById("view");
//    var nextBufferSeg = document.getElementById("next-buffer-seg");
    var buffer = document.getElementById("buffer");
    var nextBufferSeg = buffer.querySelector(".buffer-seg");

    // Check if the elements exist
    if (view && nextBufferSeg && buffer) {
        // Move the next buffer frame to the view
        view.appendChild(nextBufferSeg);

        requestAnimationFrame(() => {
            nextBufferSeg.classList.remove("hidden");
        });

        var oldSeg = view.querySelector(".seg:not(.hidden)");
        if (oldSeg) {
            oldSeg.classList.add("hidden");
            setTimeout(() => {
                if (oldSeg.parentNode) {
                    oldSeg.parentNode.removeChild(oldSeg);
                }
            }, (1000 * 2));
        }

        nextBufferSeg.removeAttribute("id");

        // Set the ID on the next buffer frame
        var bufferSeg = buffer.querySelector(".buffer-seg:not([id])");
        if (bufferSeg) {
            bufferSeg.setAttribute("id", "next-buffer-seg");
        }
    } else {
        console.error("One or more elements not found:", { view, nextBufferSeg, buffer });
    }

    // Schedule the next update
//    setTimeout(updateSeg, 1000 * 2.5);
//    nextBufferSeg.removeEventListener("ended", func);
    nextBufferSeg.onended = (event) => {
        updateSeg();
    }
}

// Start the update loop
//setTimeout(updateSeg, 10);
setTimeout(() => {
    document.getElementById("view").querySelector(".seg").onended = (event) => {
        updateSeg();
    }}, 10);

