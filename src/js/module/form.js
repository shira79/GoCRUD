'use strict';

// DOM Tree の構築が完了したら処理を開始します。
document.addEventListener('DOMContentLoaded', () => {

  // DOM API を利用して HTML 要素を取得します。
  const inputs = document.getElementsByTagName('input');
  const form = document.forms.namedItem('article-form');
  const saveBtn = document.querySelector('.article-form__save');
  const cancelBtn = document.querySelector('.article-form__cancel');

  const errors = document.querySelector('.article-form__errors');
  const errorTmpl = document.querySelector('.article-form__error-tmpl').firstElementChild;

  // 新規作成画面か編集画面かを URL から判定します。
  const mode = { method: '', url: '' };
  if (window.location.pathname.endsWith('new')) {
    mode.method = 'POST';
    mode.url = '/articles';
  } else if (window.location.pathname.endsWith('edit')) {
    mode.method = 'PATCH';
    mode.url = `/articles/${window.location.pathname.split('/')[2]}`;
  }
  const { method, url } = mode;

  const csrfToken = document.getElementsByName('csrf')[0].content;

  for (let elm of inputs) {
    elm.addEventListener('keydown', event => {
      if (event.keyCode && event.keyCode === 13) {
        // キーが押された際のデフォルトの挙動をキャンセルします。
        event.preventDefault();
        return false;
      }
    });
  }

  // 前のページに戻るイベントを設定します。
  cancelBtn.addEventListener('click', event => {
    event.preventDefault();
    window.location.href = url;
  });

  // 保存処理を実行するイベントを設定します。
  saveBtn.addEventListener('click', event => {
    event.preventDefault();

    // 前回のバリデーションエラーの表示が残っている場合は削除します。
    errors.innerHTML = null;

    // フォームに入力された内容を取得します。
    const fd = new FormData(form);

    let status;

    fetch(`/api${url}`, {
      method: method,
      headers: { 'X-CSRF-Token': csrfToken },
      body: fd
    })
      .then(res => {
        status = res.status;
        return res.json();
      })
      .then(body => {
        console.log(JSON.stringify(body));

        if (status === 200) {
          window.location.href = url;
        }

        // バリデーションエラーの表示
        if (status === 422 && body.ValidationErrors) {
          showErrors(body.ValidationErrors);
        }
      })
      .catch(err => console.error(err));
  });

  // バリデーションエラーを表示する関数
  const showErrors = messages => {
    if (Array.isArray(messages) && messages.length != 0) {
      // 複数メッセージを格納するためのフラグメントを作成します。
      const fragment = document.createDocumentFragment();

      messages.forEach(message => {
        const frag = document.createDocumentFragment();
        frag.appendChild(errorTmpl.cloneNode(true));
        frag.querySelector('.article-form__error').innerHTML = message;
        fragment.appendChild(frag);
      });
      errors.appendChild(fragment);
    }
  };

});