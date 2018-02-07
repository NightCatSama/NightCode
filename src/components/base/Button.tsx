import * as React from 'react'

export interface ButtonProps {
  onClick?: React.FormEventHandler<any>
  children: React.ReactChild
}

export default (props: ButtonProps): JSX.Element => {
  const {
    onClick,
    children
  } = props
  return (
    <button
      onClick={onClick}
    >
      {
        children
      }
    </button>
  )
}