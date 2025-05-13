import { useState } from "react";
import { Button } from "./components/ui/button";
import { ThemeProvider } from "./components/theme-provider";

function App() {
  const [count, setCount] = useState(0);

  return (
    <ThemeProvider>
      <Button onClick={() => setCount(count + 1)}>{count}</Button>
    </ThemeProvider>
  );
}

export default App;
