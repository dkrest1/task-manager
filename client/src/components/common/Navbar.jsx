import { useState } from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
  const [isLoggedIn, setLoggedIn] = useState(false);

  const handleLogout = () => {
    // Implement your logout logic and update the state
    setLoggedIn(false);
  };

  return (
    <nav className="bg-blue-500 p-4">
      <div className="container mx-auto flex items-center justify-between">
        <Link to="/" className="text-white text-2xl font-bold">
          TASKY
        </Link>
        <div className="space-x-4 flex items-center">
          <div className="lg:hidden">
            <button
              className="text-white focus:outline-none"
              // Toggle mobile menu visibility
              // You can implement this toggle functionality using state if needed
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M4 6h16M4 12h16m-7 6h7"
                />
              </svg>
            </button>
          </div>
          <div className="hidden lg:flex space-x-4">
            <Link to="/" className="text-white hover:underline">
              Home
            </Link>
            {!isLoggedIn && (
              <>
                <Link to="/signup" className="text-white hover:underline">
                  Sign Up
                </Link>
                <Link to="/login" className="text-white hover:underline">
                  Login
                </Link>
              </>
            )}
            {isLoggedIn && (
              <>
                <Link to="/profiles" className="text-white hover:underline">
                  Profiles
                </Link>
                <Link to="/tasks" className="text-white hover:underline">
                  Tasks
                </Link>
                <button
                  className="text-white hover:underline"
                  onClick={handleLogout}
                >
                  Logout
                </button>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
