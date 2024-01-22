import PropTypes from "prop-types";

const Button = ({ onClick, className, children }) => {
  return (
    <button
      onClick={onClick}
      className={`bg-primary text-white py-2 px-4 rounded-md hover:bg-blue-700 ${className}`}
    >
      {children}
    </button>
  );
};

Button.propTypes = {
  onClick: PropTypes.func,
  className: PropTypes.string,
  children: PropTypes.node.isRequired,
};

export default Button;
