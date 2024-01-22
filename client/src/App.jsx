import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Landing from "./components/pages/Landing";
import Login from "./components/pages/Login";
import SignUp from "./components/pages/SignUp";
import Navbar from "./components/common/Navbar";
import Footer from "./components/common/Footer";
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

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
     <ToastContainer />
    </Router>
  )
}

export default App
