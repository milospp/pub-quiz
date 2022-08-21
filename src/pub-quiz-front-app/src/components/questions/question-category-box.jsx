import React from 'react'

import './question-category.style.css'

export default function QuestionCategoryBox(props) {
    const {category} = props 
  return (
    <div className='category-option'>
        <h5>{category.title}</h5>
    </div>
  )
}
