import React from 'react'
import QuestionCategoryBox from './question-category-box'

export default function QuestionCategories() {

  const category = {title: 'Test'}

  return (
    <div>
        <QuestionCategoryBox category={category} />
    </div>
  )
}
