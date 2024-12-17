import React from 'react';
import ReactDOM from 'react-dom/client';
import  App  from './app/App';

// Получаем точку вмонтирования нашего приложения
const root = ReactDOM.createRoot(document.getElementById('root'));

// Вызываем метод вмонтирования приложения
root.render(
  <React.StrictMode>
    {/* // Рисуем компоненту приложения */}
    <App />
  </React.StrictMode>
);
