const serverAddr = import.meta.env.VITE_SERVER;

function App() {
    function clickHandler() {
        const req = fetch(`${serverAddr}`);

        const msg = req.then((res) => {
            console.log(res.statusText);
            return res.json();
        });

        msg.then((m) => {
            console.log(m.message);
        });
    }

    return (
        <>
            <div className="text-blue-800">Hello, wishy</div>
            <button
                className="p-2 rounded-2xl text-stone-400 bg-stone-900"
                onClick={clickHandler}
            >
                Click me to send a request to the server
            </button>
        </>
    );
}

export default App;
