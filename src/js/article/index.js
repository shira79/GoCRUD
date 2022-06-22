'use strict';

document.addEventListener('DOMContentLoaded', () => {

  const deleteBtns = document.querySelectorAll('.articles__item-delete');
  const moreBtn = document.querySelector('.page__more');
  const articles = document.querySelector('.articles');
  const articleTmpl = document.querySelector('.articles__item-tmpl').firstElementChild;

  const csrfToken = document.getElementsByName('csrf')[0].content;

  // 記事を削除する関数を定義します。
  const deleteArticle = id => {
    let statusCode;

    // Fetch API を利用して削除リクエストを送信します。
    fetch(`/api/articles/${id}`, {
      method: 'DELETE',
      headers: { 'X-CSRF-Token': csrfToken }
    })
      .then(res => {
        statusCode = res.status;
        return res.json();
      })
      .then(data => {
        console.log(JSON.stringify(data));
        if (statusCode == 200) {
          // 削除に成功したら画面から記事の HTML 要素を削除します。
          document.querySelector(`.articles__item-${id}`).remove();
        }
      })
      .catch(err => console.error(err));
  };

  // 削除ボタンそれぞれに対してイベントリスナーを設定します。
  for (let elm of deleteBtns) {
    elm.addEventListener('click', event => {
      event.preventDefault();

      // 削除ボタンのカスタムデータ属性からIDを取得して引数に渡します。
      deleteArticle(elm.dataset.id);
    });
  }

    // もっとみるボタンにイベントリスナーを設定します。
    moreBtn.addEventListener('click', event => {
        event.preventDefault();

        const cursor = moreBtn.dataset.cursor;

        if (!cursor || cursor <= 0) {
            moreBtn.remove();
            return;
        }
        // Fetch API を利用して非同期リクエストを実行します。
        let statusCode;
        fetch(`/api/articles?cursor=${cursor}`)
        .then(res => {
            statusCode = res.status;
            return res.json();
        })
        .then(data => {
            console.log(JSON.stringify(data));
            // リクエストに成功し、記事一覧データが配列で返ってきた場合
            if (statusCode == 200 && Array.isArray(data)) {
            // 表示する記事がこれ以上存在しない場合は、もっとみるボタンを画面から削除して処理を終了します。
            if (data.length == 0) {
                moreBtn.remove();
                return;
            }

            // 記事の HTML 要素をまとめるためのフラグメントを作成します。（記事のリスト）
            const fragment = document.createDocumentFragment();

            
            data.forEach(article => {
                // 個々の記事の HTML 要素を格納するフラグメントを作成します。（個別記事）
                const frag = document.createDocumentFragment();

                // 記事の HTML 要素のテンプレートからクローンを作成し、フラグメントの子要素として追加します。
                frag.appendChild(articleTmpl.cloneNode(true));

                frag.querySelector('article').classList.add(`articles__item-${article.id}`);
                frag.querySelector('.articles__item').setAttribute('href', `/articles/${article.id}`);
                frag.querySelector('.articles__item-title').textContent = article.title;
                frag.querySelector('.articles__item-date').textContent = article.created.split('T')[0];

                const deleteBtnElm = frag.querySelector('.articles__item-delete');
                deleteBtnElm.dataset.id = article.id;
                deleteBtnElm.addEventListener('click', event => {
                event.preventDefault();
                deleteArticle(article.id);
                });

                fragment.appendChild(frag);
            });

            moreBtn.dataset.cursor = data[data.length - 1].id;

            articles.appendChild(fragment);
            }
        })
        .catch(err => console.error(err));
    });
});