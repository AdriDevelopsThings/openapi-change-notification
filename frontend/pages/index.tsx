import { useState } from 'react'
import Swal from 'sweetalert2'
import styles from '../styles/index.module.css'
import errorSwal from '../utils/errorSwal'

export default function Home() {
  const [email, setEmail] = useState('')
  const [openApiUrl, setOpenApiUrl] = useState('')
  const [openApiPath, SetOpenApiPath] = useState('')
  const [method, setMethod] = useState('GET')

  return (
    <div className='container'>
      <h1>OpenAPI change notification</h1>
      <p>Receive notifications when an endpoint of an OpenAPI configuration file gets deprecated.</p>
      <form onSubmit={(e) => {
        let obj = {
          email,
          openapi_url: openApiUrl,
          path: openApiPath,
          method
        }
        Swal.showLoading(null)
        fetch(process.env.NEXT_PUBLIC_API_BASE + '/subscribe', {
          method: 'POST',
          headers: { "Content-Type": "application/json"},
          body: JSON.stringify(obj),
        })
          .then(r => r.json())
          .then(j => {
            if (Object.hasOwn(j, 'error')) {
              errorSwal(j.error)
            } else {
              Swal.fire('Success', 'You successfuly subscribed to this openapi path. If you are using this service for the first time check your mailbox for a verification email.', 'success')
            }
          }).catch(e => {
            errorSwal('default')
          })
        console.log(obj)
        e.preventDefault()
      }}>
        <div className={'form-group ' + styles.input_field}>
          <label htmlFor='email-input'>The mail address the notifications will be sent to</label>
          <input
            type="email"
            value={email}
            onChange={(e => setEmail(e.target.value))}
            id='email-input'
            className='form-control'
            placeholder='Your mail address'
           />
        </div>
        <div className={'form-group ' + styles.input_field}>
          <label htmlFor='openapiurl-input'>The url of the OpenAPI configuration (json)</label>
          <input
            type="text"
            value={openApiUrl}
            onChange={(e => setOpenApiUrl(e.target.value))}
            id='openapiurl-input'
            className='form-control'
            placeholder='OpenAPI configuration url'
            aria-describedby='openapiurl-help'
           />
           <small id="openapiurl-help" className='form-text text-muted'>For example: https://docs.bahn.expert/swagger.json</small>
        </div>
        <div className={'form-group ' + styles.input_field}>
          <label htmlFor='openapipath-input'>The path you want to subscribe to</label>
          <input
            type="text"
            value={openApiPath}
            onChange={(e => SetOpenApiPath(e.target.value))}
            id='openapipath-input'
            className='form-control'
            placeholder='OpenAPI path'
            aria-describedby='openapipath-help'
           />
           <small id="openapipath-help" className='form-text text-muted'>For example: https://bahn.expert/api/hafas/v2/arrivalStationBoard</small>
        </div>
        <div className={'form-group ' + styles.input_field}>
          <label htmlFor='openapimethod-input'>The method of the path you want to subscribe to</label>
          <select className='form-control' id="openapimethod-input" onChange={(e) => {console.log(e); setMethod(e.target.value)}} value={method}>
            <option>GET</option>
            <option>POST</option>
            <option>PUT</option>
            <option>PATCH</option>
            <option>DELETE</option>
            <option>HEAD</option>
            <option>OPTIONS</option>
            <option>CONNECT</option>
            <option>TRACE</option>
            <option>PATCH</option>
          </select>
        </div>
        <button type="submit" className={"btn btn-primary " + styles.input_field}>Submit</button>
      </form>
    </div>
  )
}
