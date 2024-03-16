import React from 'react'

function index({send}) {
    function handleKeyPress(event) {
        if (event.key === 'Enter') {
          send(event);
          event.target.value = '';
        }
      }
  return (
    <div className='m-4 border w-80'>
        <input onKeyPress={handleKeyPress} type="text" placeholder='Enter a Message' />
    </div>
  )
}

export default index