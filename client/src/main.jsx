import { createRoot } from 'react-dom';
import App from './App';
import './index.css'
import { Provider } from 'react-redux';
import { store } from './components/redux/store';



const root = createRoot(document.getElementById('root'));

root.render(
  <Provider store={store}>
    <App />
  </Provider>
);
