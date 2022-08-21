import {Routes, Route} from 'react-router-dom';
import './App.css';
import './assets/style/global.style.scss'
import './assets/style/util.style.scss'
import './assets/style/comp.style.scss'

import { Home } from './routes/home/home.component';
import { HomeTemplate } from './routes/home-template/home-template.component';
import { QuizTemplate } from './routes/quiz-template/quiz-template.component'
import { Questions } from './routes/quiz/questions.component'
import { QuizCreator } from './routes/quiz/quiz-creator.component'
import { Login } from './routes/login/login.component'
import QuizLobby from './routes/quiz-app/quiz-lobby/quiz-lobby.component';
import QuizGamePlayer from './routes/quiz-app/quiz-game-player/quiz-game-player.component';

function App() {
  return (
    <Routes>
      <Route path='/' element={<HomeTemplate />}>
        <Route path='/login' element={<Login />} /> 
        <Route path='/questions' element={<Questions />} /> 
        <Route path='/create-quiz' element={<QuizCreator />} /> 

      </Route> 
      <Route path='/' element={<QuizTemplate />}> 

        <Route path='/lobby' element={<QuizLobby />} /> 
        <Route path='/player' element={<QuizGamePlayer />} /> 

      </Route>

    </Routes>
  );
}

export default App;
