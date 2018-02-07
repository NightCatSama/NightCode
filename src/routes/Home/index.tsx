import * as React from 'react'
import { connect } from 'react-redux'

import classNames from 'classnames/bind'
import styles from './style.scss'

let cx = classNames.bind(styles)

class Home extends React.Component {
  render(): JSX.Element {
    return (
      <div className={cx('container')}>
        <h1 className={cx('logo')}>Test</h1>
      </div>
    )
  }  
}

export default connect(
  state => ({
    state
  })
)(Home as React.ComponentClass<any>)
