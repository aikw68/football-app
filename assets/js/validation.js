document.addEventListener('DOMContentLoaded', () => {
    //.validationForm を指定した最初の form 要素を取得
    const validationForm = document.querySelector('.validationForm');
    //.validationForm を指定した form 要素が存在すれば
    if(validationForm) {
      //エラーを表示する span 要素に付与するクラス名（エラー用のクラス）
      const errorClassName = 'error';
      
      //email クラスを指定された要素の集まり
      const emailElems =  document.querySelectorAll('.email');
      //password クラスを指定された要素の集まり
      const passwordElems =  document.querySelectorAll('.password');
      //equal-to クラスを指定された要素の集まり
      const equalToElems = document.querySelectorAll('.equal-to'); 
      
      //エラーメッセージを表示する span 要素を生成して親要素に追加する関数
      //elem ：対象の要素
      //errorMessage ：表示するエラーメッセージ
      const createError = (elem, errorMessage) => {
        //span 要素を生成
        const errorSpan = document.createElement('span');
        //エラー用のクラスを追加（設定）
        errorSpan.classList.add(errorClassName);
        //aria-live 属性を設定
        errorSpan.setAttribute('aria-live', 'polite');
        //引数に指定されたエラーメッセージを設定
        errorSpan.textContent = errorMessage;
        //elem の親要素の子要素として追加
        elem.parentNode.appendChild(errorSpan);
      }
   
      //form 要素の submit イベントを使った送信時の処理
      validationForm.addEventListener('submit', (e) => {
        //エラーを表示する要素を全て取得して削除（初期化）
        const errorElems = validationForm.querySelectorAll('.' + errorClassName);
        errorElems.forEach( (elem) => {
          elem.remove(); 
        });
        
        //.email を指定した要素を検証
        emailElems.forEach( (elem) => {
          //Email の検証に使用する正規表現パターン
          const pattern = /^([a-z0-9\+_\-]+)(\.[a-z0-9\+_\-]+)*@([a-z0-9\-]+\.)+[a-z]{2,6}$/ui;
          //値が空でなければ
          if(elem.value !=='') {
            //test() メソッドで値を判定し、マッチしなければエラーを表示してフォームの送信を中止
            if(!pattern.test(elem.value)) {
              createError(elem, '正しいメールアドレス形式で入力してください。');
              e.preventDefault();
            }
          }
        });

        //.password を指定した要素を検証
        passwordElems.forEach( (elem) => {
          //password の検証に使用する正規表現パターン（8文字以上、半角英数字を必ず使用すること）
          const patternPassword = /^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-z0-9]{8,}$/ui;
          //値が空でなければ
          if(elem.value !=='') {
            //test() メソッドで値を判定し、マッチしなければエラーを表示してフォームの送信を中止
            if(!patternPassword.test(elem.value)) {
              createError(elem, '正しいパスワード形式で入力してください。');
              e.preventDefault();
            }
          }
        });

        //.equal-to を指定した要素を検証
        equalToElems.forEach( (elem) => {
          //比較対象の要素の id 
          const equalToId = elem.dataset.equalTo;
          //または const equalToId = elem.getAttribute('data-equal-to');
          //比較対象の要素
          const equalToElem = document.getElementById(equalToId);
          //値が空でなければ
          if(elem.value !=='' && equalToElem.value !==''){
            if(equalToElem.value !== elem.value) {
              createError(elem, '入力されたパスワードが一致しません。');
              e.preventDefault();
            }
          }
        });
   
        //エラーの最初の要素を取得
        const errorElem =  validationForm.querySelector('.' + errorClassName);
        //エラーがあればエラーの最初の要素の位置へスクロール
        if(errorElem) {
          const errorElemOffsetTop = errorElem.offsetTop;
          window.scrollTo({
            top: errorElemOffsetTop - 40,  //40px 上に位置を調整
            //スムーススクロール
            behavior: 'smooth'
          });
        }
      }); 
    }
  });