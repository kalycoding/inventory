import './Category.css';
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
  console.log(categoryData)
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
      <h1 class="title">Category</h1>
      <div class="container">
        <form onSubmit={handleSubmit}>
          <div class="row">
            <div class="col-25">
              <label for="fname">Category Name</label>
            </div>
            <div class="col-75">
              <input
                type="text"
                name="category_name"
                placeholder="Category Name"
                value={formValue.category_name}
                onChange={handleChange}
                required
              />
              </div>
          </div>
          <div class="row">
            <div class="col-25">
              <label for="fname">Category Description</label>
            </div>
            <div class="col-75">
              <textarea
                type="text"
                name="category_description"
                placeholder="Category Descriptiom"
                value={formValue.category_description}
                onChange={handleChange}
                required
              />
              </div>
          </div>
          <div class="row">
            <input type="submit" value="Submit"/>
          </div>
        </form>
      </div>  
      <div>
        <p>
          {
            !categoryData ? <p>Empty Categories, Please Add</p> : <ul>{categoryData.map(data=>(<span><li key={data.id}>{data.name}</li><p><button class="delete" onClick={()=>handleDelete(data.id)}>Delete</button></p></span>))}</ul>
          } 
        </p>
      </div>
    </div>
  )
}

export default App;


/* <div class="container">
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
      <textarea
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
        Category
      </button>
      <p>Category</p> 
    </form>
    {
      !categoryData ? <p>No data</p> : <ul>{categoryData.map(data=>(<button key={data.id} onClick={()=>handleDelete(data.id)}>{data.name}</button>))}</ul>
    }
    </div> */