import './App.css';
import React from 'react';
import axios from 'axios';
function App() {

  const [formValue, setformValue] = React.useState({
    category_name: '',
    category_description: ''
  });

  const handleSubmit = (event) => {
    const categoryFormData = new FormData();
    categoryFormData.append("username", formValue.category_name)
    categoryFormData.append("category_description", formValue.category_description)
    alert(formValue.category_name)
    const article = {name:formValue.category_name,description:formValue.category_description};
      axios.post('http://localhost:8080/api/category',article)
          .then(response => alert(response.data))
          .catch(error => {
            alert(error)
          })
  
  }

  const handleChange = (event) => {
    setformValue({
      ...formValue,
      [event.target.name]: event.target.value
    });
  }
  

  return(
    <form onSubmit={handleSubmit}>
      <p>Category Creation Page</p>
      <input
        type="text"
        name="category_name"
        placeholder="Category Name"
        value={formValue.category_name}
        onChange={handleChange}
      />
      <input
        type="text"
        name="category_description"
        placeholder="Category Description"
        value={formValue.category_description}
        onChange={handleChange}
      />
      <button
        type="submit"
      >
        Login
      </button>
    </form>
  )
}

export default App;

