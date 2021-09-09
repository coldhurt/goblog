import React from 'react'
import loading from './loading.svg'

const Spin = () => {
  return (
    <img
      src={loading}
      style={{
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: '-50% -50%',
        height: 100,
      }}
    />
  )
}

export default Spin
