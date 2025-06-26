import { BrowserRouter, Route, Routes } from 'react-router';

// const serverAddr = import.meta.env.VITE_SERVER;

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<div>Hello, world</div>}></Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
