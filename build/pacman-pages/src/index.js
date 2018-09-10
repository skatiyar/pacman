import './base.scss';
import './styles.scss';

const loading = document.getElementById('pacman-game-loading');
const iframe = document.getElementById('pacman-game');
iframe.onload = () => {
    iframe.className = 'PacmanGame';
    loading.className = 'PacmanGameLoading hide';
};
