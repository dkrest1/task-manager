import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Landing from "./components/pages/Landing";
import Login from "./components/pages/Login";
import SignUp from "./components/pages/SignUp";
import Navbar from "./components/common/Navbar";
import Footer from "./components/common/Footer";

function App() {
  return (
    <Router>
      <Navbar/>
     <Routes>
      <Route path="/" exact Component={Landing} />
      <Route path="/login" Component={Login}/>
      <Route path="/signup" Component={SignUp}/>
     </Routes>
     <Footer/>
    </Router>
  )
}

export default App
