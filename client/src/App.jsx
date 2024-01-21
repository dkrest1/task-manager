import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Landing from "./components/pages/Landing";
import Login from "./components/pages/Login";
import SignUp from "./components/pages/SignUp";

function App() {
  return (
    <Router>
     <Routes>
      <Route path="/" exact Component={Landing} />
      <Route path="/login" Component={Login}/>
      <Route path="/signup" Component={SignUp}/>
     </Routes>
    </Router>
  )
}

export default App
