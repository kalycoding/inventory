import './App.css';
import React from 'react';
import axios from 'axios';


function ProductApp() {
  const [formValue, setformValue] = React.useState({
    category_id:'',
    product_name: '',
    product_description: ''
  });  

  const handleDelete = (data)=> {
    axios.delete(`http://localhost:8080/api/product/${data}`)
          .then(function(response){
         
          })
          window.location.reload(); 
  }

  const [productData, setProduct] = React.useState([]);
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
  React.useEffect(() => {
    let mounted = true;
    axios.get('http://localhost:8080/api/product')
          .then(function(response){
            if (response.status===200){
              if(mounted) {
                  setProduct(response.data)
                }
            }
            
          })
          .catch(error => {setProduct([])
          })
    
    // getList()
    //   .then(items => {
    //     
    //   })
    return () => mounted = false;
  }, [])

  console.log(productData)
  const handleSubmit = (event) => {
    //event.preventDefault();
    const productFormData = new FormData();
    productFormData.append("category_id", formValue.category_id)
    productFormData.append("username", formValue.product_name)
    productFormData.append("category_description", formValue.product_description)
    const body = {name:formValue.product_name,description:formValue.product_description};
      axios.post('http://localhost:8080/api/product',body, { params: { catId: formValue.category_id } })
      .then(function(response){
        if (response.status===201){
          alert(`Product ${formValue.product_name} Created`)
        }
        else{
          alert(`Product ${formValue.product_name} Already Exist`)
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
      <h1 class="title">Product</h1>
        <div class="container">
          <form onSubmit={handleSubmit}>
          
            <div class="row">
              <div class="col-25">
                <label for="fname">Category</label>
              </div>
              <div class="col-75">
                {
                  !categoryData ? <select><option value="none" selected disabled hidden >Empty Categories, Add Category</option></select> : 
                  <select name="category_id" onChange={handleChange}>
                    <option value="none" selected disabled hidden>Choose Category</option>
                    {categoryData.map(data=>(
                    <option key={data.id} value={data.id} name="product_id">{data.name}</option>
                    ))}
                  </select>
                }
              </div>
              <div class="col-25">
                <label for="fname">Product Name</label>
              </div>
              <div class="col-75">
                <input
                  type="text"
                  name="product_name"
                  placeholder="Product Name"
                  value={formValue.product_name}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>
            <div class="row">
              <div class="col-25">
                <label for="fname">Product Description</label>
              </div>
              <div class="col-75">
                <input
                  type="text"
                  name="product_description"
                  placeholder="Product Description"
                  value={formValue.product_description}
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
            !productData ? <p>Empty Product, Please Add</p> : <ul>{productData.map(data=>(<span><li key={data.id}>{data.name}</li><p><button class="delete" onClick={()=>handleDelete(data.id)}>Delete</button></p></span>))}</ul>
          } 
        </p>
      </div>
    </div>
    )
}

export default ProductApp;

/* <div>
        <form onSubmit={handleSubmit}>
          <p>Product Page</p>
          {
                  !categoryData ? <select><option value="none" selected disabled hidden >Empty Categories, Add Category</option></select> : 
                  <select name="category_id" onChange={handleChange}>
                    <option value="none" selected disabled hidden>Choose Category</option>
                    {categoryData.map(data=>(
                    <option key={data.id} value={data.id} name="product_id">{data.name}</option>
                    
                    ))}
                  </select>
                }
          <input
            type="text"
            name="product_name"
            placeholder="Product Name"
            value={formValue.product_name}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="product_description"
            placeholder="Product Description"
            value={formValue.product_description}
            onChange={handleChange}
            required
          />
          <button
            type="submit"
          >
            Create Product
          </button>
          <p>Product List</p> 
        </form>
        {
        !productData ? <p>No data</p> : 
        <ul>
            {productData.map(data=>(<button key={data.id} 
              onClick={()=>handleDelete(data.id)}>
                {data.name}</button>)
            )}
        </ul>
        }
        </div> */