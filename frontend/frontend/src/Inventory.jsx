import './App.css';
import React from 'react';
import axios from 'axios';



function InventoryApp(){
    const [formValue, setformValue] = React.useState({
        product_id:'',
        cost_price: 0,
        quantity: 0,
        stock_level: 0,
        selling_price:0,
        supplier:''
      });  
    
      const handleDelete = (data)=> {
        axios.delete(`http://localhost:8080/api/inventory/${data}`)
              .then(function(response){
             
              })
              window.location.reload(); 
      }
    
    
      const [inventoryData, setInventory] = React.useState([]);
      const [productData, setProduct] = React.useState([]);
      
      React.useEffect(() => {
        let mounted = true;
        axios.get('http://localhost:8080/api/inventory')
              .then(function(response){
                if (response.status===200){
                  if(mounted) {
                      setInventory(response.data)
                    }
                }
                
              })
              .catch(error => {setInventory([])
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
    
      console.log(formValue)
      const handleSubmit = (event) => {
        //event.preventDefault();
        const productFormData = new FormData();
        productFormData.append("products", formValue.product_id)
        productFormData.append("cost_price", formValue.cost_price)
        productFormData.append("quantity", formValue.quantity)
        productFormData.append("stock_level", formValue.stock_level)
        productFormData.append("selling_price", formValue.selling_price)
        productFormData.append("supplier", formValue.supplier)
        const body = {
            costPrice : Number(formValue.cost_price), 
            quantity : Number(formValue.quantity),
            stockLevel : Number(formValue.stock_level),
            sellingPrice : Number(formValue.selling_price),
            supplier : formValue.supplier
        };
          axios.post('http://localhost:8080/api/inventory',body, { params: { prodId: formValue.product_id } })
              .then(function(response){
                // if (response.status===201){
                //   alert(`Category ${formValue.supplier} Created`)
                // }
                // else{
                //   alert(`Category ${formValue.supplier} Already Exist`)
                // }
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
            <form onSubmit={handleSubmit} id='products'>
              <p>Inventory Creation Page</p>
              

                {
                  !productData ? <select><option value="none" selected disabled hidden>Empty Products, Add product</option></select> : 
                  <select name="product_id" onChange={handleChange}>
                    <option value="none" selected disabled hidden>Choose Product</option>
                    {productData.map(data=>(
                    <option key={data.id} value={data.id} name="product_id">{data.name}</option>
                    
                    ))}
                  </select>
                }
              <input
                type="text"
                name="cost_price"
                placeholder="Cost Price"
                value={formValue.cost_price}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="quantity"
                placeholder="Quantity"
                value={formValue.quantity}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="stock_level"
                placeholder="Stock Level"
                value={formValue.stock_level}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="selling_price"
                placeholder="Selling Price"
                value={formValue.selling_price}
                onChange={handleChange}
                required
              />
              <input
                type="text"
                name="supplier"
                placeholder="Supplier"
                value={formValue.supplier}
                onChange={handleChange}
                required
              />
              <button
                type="submit"
              >
                Create Inventory
              </button>
              <p>Inventory List</p> 
            </form>
            {
            !inventoryData ? <p>No data</p> : 
            <ul>
                {inventoryData.map(data=>(<button key={data.id} 
                  onClick={()=>handleDelete(data.id)}>
                    {data.product.name}</button>)
                )}
            </ul>
            }
            </div>
          )
}

export default InventoryApp