import { toast } from "vue3-toastify";
import "vue3-toastify/dist/index.css";

const showNotification = (message: string, type: string) => {
  if (type === "success") {
    toast(message, {
      theme: "auto",
      type: "success",
      dangerouslyHTMLString: true,
      autoClose: 1000,
      position: "top-right",
    });
  } else if (type === "error") {
    toast(message, {
      theme: "auto",
      type: "error",
      dangerouslyHTMLString: true,
      autoClose: 1000,
      position: "top-right",
    });
  }
};

export default showNotification;
