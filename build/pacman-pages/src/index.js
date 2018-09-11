import './base.scss';
import './styles.scss';

const loading = document.getElementById('pacman-game-loading');
const iframe = document.getElementById('pacman-game');
iframe.onload = () => {
    loading.className = 'PacmanGameLoading hide';
    iframe.className = 'PacmanGame';
};
