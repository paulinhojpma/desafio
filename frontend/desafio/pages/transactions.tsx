

import React, { Component } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import 'bootstrap/dist/css/bootstrap.css';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

 class PageWithJSbasedForm extends Component <any, any>{
constructor(props:any) {
  super(props);
  this.state = {
    kleb: "ronaldo",
    listState: [],
    fetched: false,
  }
  this.handleSubmit = this.handleSubmit.bind(this);
}
  
   handleSubmit = async (event: any) => {
    
  event.preventDefault()
 
  this.setState({kleb: "vicente"})
   
    var d = new Date("2022-10-20T00:00:00Z")
     console.log(new Intl.DateTimeFormat('en-US', {year: 'numeric', month: '2-digit',day: '2-digit', hour: '2-digit', minute: '2-digit', second: '2-digit'}).format(d));

     console.log("DATE - ",d.toLocaleDateString())
    
     let formData = new FormData();
     formData.append("file", event.target.file.files[0])
  
    const endpoint = 'http://localhost:8890/transactions'

    var options = {

      method: 'POST',
    
      body: formData,
    }


    const response = await fetch(endpoint, options)

    if (response.status == 201) {
       console.log("Status", response.status)
    const result = await response.json()
    
    console.log(result.codResponse)
   
    var options1 = {
    method: 'GET',
   }

   const response1 = await fetch(endpoint, options1)

   if (response1.status == 200) {
    const result1 = await response1.json()
   
   var lista = []
   lista = result1.producers   
 
   this.setState({listState: lista, fetched: true})
   }else{
      const result1 = await response1.json()
    alert(`Transaction incomplete: ${result1.message}`)
   }
   
  
    }else{
       const result = await response.json()
      alert(`Transaction incomplete: ${result.message}`)
    }
   
 
    
  }


 render(): React.ReactNode {
   return (<Container>     
   
    <Form onSubmit={e => this.handleSubmit(e)}>
    <Form.Group className="mb-3" controlId="formBasicEmail">
      <Form.Label >Choose file to upload</Form.Label>
       <Form.Control type="file" name="file" placeholder="Password" />       
    </Form.Group>
    <Button variant="primary" type="submit">
        Submit
      </Button>  
    </Form>
    <br/>
    
    <Container>
    { 
      this.state.fetched &&  
      <Row className='bg-secondary'> 
        <Col  >Name</Col> 
         <Col  >Total</Col>
          <Col   >Product</Col>
            <Col  >Data</Col>
           <Col   >Value</Col>
           <Col   >Description</Col>
      </Row> 
    }
    {this.state.listState.map((item:any, index:any) => {
      var rowColor = ''
      if (index %2 == 0) {
        rowColor ='bg-light'
      }else{
        rowColor ='bg-info'
      }
      const num = item.SumTotal
        const value = 'R$' + num.toFixed(2).replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
      return (
      <Row key={index} className={rowColor} > 
      <Col  key={item.Name}>{item.Name}</Col>
      <Col  key={item.ID} >{value}</Col>
      <Col key={item.Transactions[0].Product} >{item.Transactions[0].Product}</Col>
      <Col   >
      { item.Transactions.map((trans:any, ind:any)=> {
        const d = new Date(trans.Date)
       const formDate =  new Intl.DateTimeFormat('en-US', {day: '2-digit', month: '2-digit', year: 'numeric'}).format(d)
        return  (
      
        <div   key={ind} >{formDate}</div>
        
        ) 
      }
        
        
      )}
      </Col>
      <Col >
      { item.Transactions.map((trans:any, ind:any)=> {
        var num = trans.Value
        var value = ''
        if (trans.Type.Nature == "Saída"){
          var value = 'R$-' + num.toFixed(2).replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
        }else{
          var value = 'R$' + num.toFixed(2).replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
        }
        
        return  (
      
        <div   key={ind} >{value}</div>
        
        ) 
      }
        
        
      )}
      </Col>
       <Col   >
      { item.Transactions.map((trans:any, ind:any)=> {
        return  (
      
        <div   id={ind} key={ind} >{trans.Type.Description}</div>
        
        ) 
      }
        
        
      )}
      </Col>
       </Row> )
     
 }  )}
    </Container>
    </Container>)
 }
  
  
}

export default PageWithJSbasedForm;