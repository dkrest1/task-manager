import  { useEffect } from "react";
import PropTypes from "prop-types";
import { toast } from 'react-toastify';

export default function Alert({ type, message }) {
  useEffect(() => {
    if (type === "success") {
      toast.success(message);
    } else if (type === "error") {
      toast.error(message);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []); 

  return null;
}

Alert.propTypes = {
  type: PropTypes.oneOf(["success", "error"]).isRequired,
  message: PropTypes.string.isRequired,
};
