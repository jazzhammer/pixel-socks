const initWebsocket = () => {
  try {
    const wsInstance = new WebSocket(process.env.NEXT_PUBLIC_WEBSOCKET_URL);
    console.log(`initialized websocket`);
    return wsInstance;
  } catch (err) {
    console.log(`failed to initialize websocket:`);
    console.error(err);
  }
}
const wsInstance = initWebsocket();
wsInstance.onmessage = (data) => {
  console.log(`ws message: ${JSON.stringify(data)}`);
}
export default wsInstance;