import './App.css';
import React from 'react';
import axios from 'axios';

function App() {

  const [formValue, setformValue] = React.useState({
    category_name: '',
    category_description: ''
  });  

  const handleDelete = (data)=> {
    axios.delete(`http://localhost:8080/api/category/${data}`)
          .then(function(response){
         
          })
          window.location.reload(); 
  }

  const [categoryData, setCategory] = React.useState([]);

  React.useEffect(() => {
    let mounted = true;
    axios.get('http://localhost:8080/api/category')
          .then(function(response){
            if (response.status===200){
              if(mounted) {
                  setCategory(response.data)
                }
            }
            
          })
          .catch(error => {setCategory([])
          })
    
    // getList()
    //   .then(items => {
    //     
    //   })
    return () => mounted = false;
  }, [])
  
// try {
//   const listItems = categoryData.map((data) =>
//   <li onClick={()=>handleDelete(data.id)}>{data.name}</li>
//   );
// } catch (error) {
//   alert("Error")
// }
  console.log(categoryData)
  const handleSubmit = (event) => {
    //event.preventDefault();
    const categoryFormData = new FormData();
    categoryFormData.append("username", formValue.category_name)
    categoryFormData.append("category_description", formValue.category_description)
    const body = {name:formValue.category_name,description:formValue.category_description};
      axios.post('http://localhost:8080/api/category',body)
          .then(function(response){
            if (response.status===201){
              alert(`Category ${formValue.category_name} Created`)
            }
            else{
              alert(`Category ${formValue.category_name} Already Exist`)
            }
          })
          .catch(error => {alert(error)
          })
  
  }

  const handleChange = (event) => {
    setformValue({
      ...formValue,
      [event.target.name]: event.target.value
    });
  }
  

  return(
    <div>
    <form onSubmit={handleSubmit}>
      <p>Category Creation Page</p>
      <input
        type="text"
        name="category_name"
        placeholder="Category Name"
        value={formValue.category_name}
        onChange={handleChange}
        required
      />
      <input
        type="text"
        name="category_description"
        placeholder="Category Description"
        value={formValue.category_description}
        onChange={handleChange}
        required
      />
      <button
        type="submit"
      >
        Create Category
      </button>
      <p>Category</p> 
    </form>
     {
      !categoryData ? <p>No data</p> : <ul>{categoryData.map(data=>(<button key={data.id} onClick={()=>handleDelete(data.id)}>{data.name}</button>))}</ul>
    }
    </div>
  )
}

export default App;
