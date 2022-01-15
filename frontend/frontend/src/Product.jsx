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
              alert(`Category ${formValue.product_name} Created`)
            }
            else{
              alert(`Category ${formValue.product_name} Already Exist`)
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
            name="category_id"
            placeholder="Category"
            value={formValue.category_id}
            onChange={handleChange}
            required
          />
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
        </div>
      )
}

export default ProductApp;

