// function FormTest() {
//   return (<form action="http://localhost:8890/transactions" method="post" encType="multipart/form-data">
//         <div>
//             <label htmlFor="file">Choose file to upload</label>
//             <input type="file" id="file" name="file" multiple />
//         </div>
//         <div>
//             <button>Submit</button>
//         </div>
//     </form>)
// }

// export default FormTest


export default function PageWithJSbasedForm() {
  // Handles the submit event on form submit.
  const handleSubmit = async (event: any) => {
    // Stop the form from submitting and refreshing the page.
   event.preventDefault()

    // Get data from the form.
    const data = {
      first: event.target.file.value,
     // last: event.target.last.value,
    }
    console.log(event.target.file.files[0])
    // Send the data to the server in JSON format.
    //const JSONdata = JSON.stringify(data)
     let formData = new FormData();
     formData.append("file", event.target.file.files[0])
    // API endpoint where we send form data.
    const endpoint = 'http://localhost:8890/transactions'

    // Form the request for sending data to the server.
    var options = {
      // The method is POST because we are sending data.
      method: 'POST',
      // Tell the server we're sending JSON.
    //   headers: {
        
    //     'Content-Disposition': 'form-data'
    //   },
      // Body of the request is the JSON data we created above.
      body: formData,
    }

    // Send the form data to our forms API on Vercel and get a response.
    const response = await fetch(endpoint, options)
    console.log(response)
    // Get the response data from server as JSON.
    // If server returns the name submitted, that means the form works.
    const result = await response.json()
    console.log(JSON.stringify(result))
   // alert(`Transaction complete: ${result.codResponse}`)
    var options1 = {
    method: 'GET',
   }

   const response1 = await fetch(endpoint, options1)
   const result1 = response1.json()
   result1.producers.forEach(element => console.log(element))
   
  }
  return (
    // We pass the event to the handleSubmit() function on submit.
    <form onSubmit={handleSubmit}>
      <label htmlFor="file">Choose file to upload</label>
      <input type="file" id="file" name="file" multiple  required />


      <button type="submit">Submit</button>
    </form>
  )
}
