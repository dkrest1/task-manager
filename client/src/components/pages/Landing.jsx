import { Link } from "react-router-dom";
import Button from "../common/Button";

const Landing = () => {

  return (
    <div className="text-center mt-16">
      <h1 className="text-4xl sm:text-5xl md:text-6xl font-bold text-blue-500 mb-4">TASKY</h1>
      <p className="text-lg text-gray-600 mb-8">Manage all Your Tasks at One Go!</p>
      <div className="flex flex-col sm:flex-row justify-center">
        <Link to="/login" className="mb-4 sm:mr-4">
          <Button>Login</Button>
        </Link>
        <Link to="/signup">
          <Button>Sign Up</Button>
        </Link>
      </div>
    </div>
  );
};

export default Landing;
