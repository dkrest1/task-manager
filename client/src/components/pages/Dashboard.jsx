import { useSelector } from 'react-redux';
import { Link } from 'react-router-dom';

const Dashboard = () => {
  const { userInfo } = useSelector((state) => state.auth);

  return (
    <div className="container mx-auto mt-8">
      <h2 className="text-3xl font-bold mb-4">Welcome, {userInfo.username}!</h2>

      <div className="grid grid-cols-2 gap-4">
        <Link to="/profiles" className="p-4 bg-blue-500 text-white rounded-md">
          View Profiles
        </Link>

        <Link to="/tasks" className="p-4 bg-green-500 text-white rounded-md">
          View Tasks
        </Link>

        <Link to="/create-task" className="p-4 bg-yellow-500 text-white rounded-md">
          Create Task
        </Link>

        <Link to="/update-task" className="p-4 bg-yellow-500 text-white rounded-md">
          Update Task
        </Link>

        <Link to="/view-task/123" className="p-4 bg-yellow-500 text-white rounded-md">
          View Single Task
        </Link>

        <Link to="/update-profile" className="p-4 bg-purple-500 text-white rounded-md">
          Update Profile
        </Link>

        <button className="p-4 bg-red-500 text-white rounded-md">
          Delete Profile
        </button>

        <button className="p-4 bg-red-500 text-white rounded-md">
          Delete Task
        </button>
      </div>
    </div>
  );
};

export default Dashboard;
