import { useState } from "react";
import { Link } from "react-router-dom";
import {useDispatch, useSelector} from "react-redux"
import { loginAsync, setError } from "../../redux/authSlice";
import Alert from "../common/Alert";

const Login = () => {
  const [loginData, setLoginData] = useState({
    email: '',
    password: ''
  })

  const dispatch = useDispatch()
  const {loading, error, success} = useSelector((state) => state.auth)

  const handleSubmit = async(e) => {
    e.preventDefault()
    const {email, password} = loginData

    try {
      dispatch(loginAsync({email, password}))
      setLoginData({
        email: '',
        password: ''
      })
    }catch(error) {
      dispatch(setError(error))
    }
  }
  
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="max-w-md w-full p-6 bg-white shadow-lg rounded-md">
        <h2 className="text-3xl font-bold text-center text-primary mb-6">
          Login
        </h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Email
            </label>
            <input
              type="email"
              name="email"
              className="w-full px-3 py-2 border rounded-md"
              placeholder="Enter your email"
              value={loginData.email}
              onChange={(e) => setLoginData({...loginData, email: e.target.value})}
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              Password
            </label>
            <input
              type="password"
              name="password"
              className="w-full px-3 py-2 border rounded-md"
              placeholder="Enter your password"
              value={loginData.password}
              onChange={(e) => setLoginData({...loginData, password: e.target.value})}
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-primary text-white py-2 rounded-md hover:bg-blue-700"
            disabled={loading}
          >
            {loading ? "Logining In..." : "Login"}
          </button>
        </form>
        {error && (
          <Alert type="error" message={error} />
        )}
        {
          success && (
            <Alert type="success" message="Login successfully" />
          )
        }
        <p className="mt-4 text-gray-600 text-center">
          Don&apos;t have an account?{" "}
          <Link to="/signup" className="text-primary hover:underline">
            Sign up here
          </Link>
        </p>
      </div>
    </div>
  );
};

export default Login;
