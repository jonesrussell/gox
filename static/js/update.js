// Create a new EventSource that connects to the /updates endpoint
const evtSource = new EventSource("/updates?stream=messages");

// Set up an event listener for the message event
evtSource.onmessage = (event) => {
  // Log the event data to the console
  console.log(event);

  // Parse the event data
  const data = event.data.split(', ');
  const title = data[0].substring(6); // Remove 'Title: ' from the start
  const body = data[1].substring(5); // Remove 'Body: ' from the start

  // Update the document's title and body
  document.title = title;
  document.body.innerHTML = body;
};

// Add an event listener for errors
evtSource.onerror = (err) => {
  console.error("EventSource failed:", err);
};
