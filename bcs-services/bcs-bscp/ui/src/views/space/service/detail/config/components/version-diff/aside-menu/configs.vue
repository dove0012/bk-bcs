<script lang="ts" setup>
  import { ref, computed, watch, onMounted } from 'vue'
  import { useRoute } from 'vue-router'
  import { storeToRefs } from 'pinia'
  import { Search, RightShape } from 'bkui-vue/lib/icon'
  import { useServiceStore } from '../../../../../../../../store/service'
  import { ICommonQuery } from '../../../../../../../../../types/index'
  import { IConfigItem, IConfigListQueryParams, IBoundTemplateGroup, IConfigDiffSelected } from '../../../../../../../../../types/config'
  import { IFileConfigContentSummary } from '../../../../../../../../../types/config';
  import { getConfigList, getConfigItemDetail, getConfigContent, getBoundTemplates, getBoundTemplatesByAppVersion } from '../../../../../../../../api/config'
  import { byteUnitConverse } from '../../../../../../../../utils'
  import SearchInput from '../../../../../../../../components/search-input.vue'

  interface IConfigMenuItem {
    type: string;
    id: number;
    name: string;
    file_type: string;
    file_state: string;
    update_at: string;
    byte_size: string;
    signature: string;
    template_revision_id: number;
  }

  interface IConfigDiffItem extends IConfigMenuItem {
    diff_type: string;
    current: string;
    base: string;
  }

  interface IConfigsGroupData {
    template_space_id: number;
    id: number;
    name: string;
    expand: boolean;
    configs: IConfigMenuItem[];
  }

  interface IDiffGroupData extends IConfigsGroupData {
    configs: IConfigDiffItem[]
  }

  const props = withDefaults(defineProps<{
    currentVersionId: number;
    baseVersionId: number|undefined;
    selectedConfig: IConfigDiffSelected;
    actived: boolean;
  }>(), {
    selectedConfig: () => ({ pkgId: 0, id: 0, version: 0 })
  })

  const emits = defineEmits(['selected'])

  const route = useRoute()
  const bkBizId = ref(String(route.params.spaceId))
  const { appData } = storeToRefs(useServiceStore())

  const diffCount = ref(0)
  const selected = ref<IConfigDiffSelected>({ pkgId: 0, id: 0, version: 0 })
  const currentGroupList = ref<IConfigsGroupData[]>([])
  const baseGroupList = ref<IConfigsGroupData[]>([])
  // 汇总的配置项列表，包含未修改、增加、删除、修改的所有配置项
  const aggregatedList = ref<IDiffGroupData[]>([])
  const groupedConfigListOnShow = ref<IDiffGroupData[]>([])
  const isOnlyShowDiff = ref(false) // 只显示差异项
  const isOpenSearch = ref(false)
  const searchStr = ref('')

  // 是否实际选择了对比的基准版本，为了区分的未命名版本id为0的情况
  const isBaseVersionExist = computed(() => {
    return typeof props.baseVersionId === 'number'
  })

  // 基准版本变化，更新选中对比项
  watch(() => props.baseVersionId, async() => {
    baseGroupList.value = await getConfigsOfVersion(props.baseVersionId)
    aggregatedList.value = calcDiff()
    groupedConfigListOnShow.value = aggregatedList.value.slice()
    setDefaultSelected()
  })

  // 当前版本默认选中的配置项
  watch(() => props.selectedConfig, (val) => {
    if (val) {
      selected.value = { ...val }
    }
  }, {
    immediate: true
  })

  onMounted(async() => {
    await getAllConfigList()
    aggregatedList.value = calcDiff()
    groupedConfigListOnShow.value = aggregatedList.value.slice()
    setDefaultSelected()
  })

  // 判断版本是否为未命名版本
  const isUnNamedVersion = (id: number) => {
    return id === 0
  }

  // 获取某一版本下配置项和模板列表
  const getConfigsOfVersion = async (releaseId: number|undefined) => {
    if (typeof releaseId !== 'number') {
      return []
    }

    const [commonConfigList, templateList] = await Promise.all([
      getCommonConfigList(releaseId),
      getBoundTemplateList(releaseId)
    ])

    return commonConfigList.concat(templateList)
  }

  // 获取非模板配置项列表
  const getCommonConfigList = async(id: number): Promise<IConfigsGroupData[]> => {
    const params: IConfigListQueryParams = {
      start: 0,
      all: true
    }
    const configsDetailQueryParams: { release_id?: number } = {}

    if (!isUnNamedVersion(id)) {
      params.release_id = id
      configsDetailQueryParams.release_id = id
    }

    const configsRes = await getConfigList(bkBizId.value, <number>appData.value.id, params)
    // 未命名版本中包含被删除的配置项，需要过滤掉
    const configs: IConfigItem[] = configsRes.details.filter((item: IConfigItem) => item.file_state !== 'DELETE')

    // 遍历配置项列表，拿到每个配置项的signature
    const configsDetailRes =  await Promise.all(configs.map(item => getConfigItemDetail(bkBizId.value, item.id, <number>appData.value.id, configsDetailQueryParams)))

    return [{
      template_space_id: 0,
      id: 0,
      name: '非配置项分组',
      expand: true,
      configs: configs.map((config, index) => {
        const { id, spec, revision, file_state } = config
        const { name, file_type } = spec
        const { byte_size, signature } = configsDetailRes[index].content
        return { type: 'config', id, name, file_type, file_state, update_at: revision.update_at, byte_size, signature, template_revision_id: 0 }
      })
    }]
  }

  // 获取模板配置项列表
  const getBoundTemplateList = async(id: number) => {
      const params: ICommonQuery = {
        start: 0,
        all: true
      }
      let res
      if (isUnNamedVersion(id)) {
        res = await getBoundTemplates(bkBizId.value, <number>appData.value.id, params)
      } else {
        res = await getBoundTemplatesByAppVersion(bkBizId.value, <number>appData.value.id, id)
      }
      return res.details.map((groupItem: IBoundTemplateGroup) => {
        const { template_space_id, template_space_name, template_set_id, template_set_name } = groupItem
        const group: IConfigsGroupData = {
          template_space_id,
          id: template_set_id,
          name: `${template_space_name === 'default_space' ? '默认空间' : template_space_name} - ${template_set_name}`,
          expand: false,
          configs: []
        }
        groupItem.template_revisions.forEach(tpl => {
          const { template_id, name, file_type, file_state, byte_size, signature, template_revision_id } = tpl
          if (file_state !== 'DELETE') {
            group.configs.push({
              type: 'template',
              id: template_id,
              name,
              file_type,
              file_state,
              update_at: '',
              byte_size,
              signature,
              template_revision_id: template_revision_id
            })
          }
        })
        return group
      })
  }

  // 获取当前版本和基准版本的所有配置项列表
  const getAllConfigList = async () => {
    currentGroupList.value = await getConfigsOfVersion(props.currentVersionId)
    baseGroupList.value = await getConfigsOfVersion(props.baseVersionId)
  }

  // 计算配置被修改、被删除、新增的差异
  const calcDiff = () => {
    diffCount.value = 0
    const list: IDiffGroupData[]= []
    currentGroupList.value.forEach(currentGroupItem => {
      const { template_space_id, id, name, expand, configs } = currentGroupItem
      const diffGroup: IDiffGroupData = { template_space_id, id, name, expand, configs: [] }
      configs.forEach(crtItem => {
        let baseItem: IConfigMenuItem|undefined
        baseGroupList.value.some(baseGroupItem => {
          if(baseGroupItem.template_space_id === currentGroupItem.template_space_id) {
            return baseGroupItem.configs.some((config) => {
              if (config.id === crtItem.id && config.template_revision_id === crtItem.template_revision_id) {
                baseItem = config
                return true
              }
            })
          }
          return false
        })
        if (baseItem) { // 修改项
            const diffConfig = {
                ...crtItem,
                diff_type: '',
                current: crtItem.signature,
                base: baseItem.signature
            }
            if (crtItem.template_revision_id !== baseItem.template_revision_id || diffConfig.current !== diffConfig.base) {
              diffCount.value++
              diffConfig.diff_type = isBaseVersionExist.value ? 'modify' : ''
            }
            diffGroup.configs.push(diffConfig)
        } else { // 当前版本新增项
            diffCount.value++
            diffGroup.configs.push({
                ...crtItem,
                diff_type: isBaseVersionExist.value ? 'add' : '',
                current: crtItem.signature,
                base: ''
            })
        }
      })
      if (diffGroup.configs.length > 0) {
        list.push(diffGroup)
      }
    })
    // 计算当前版本删除项
    baseGroupList.value.forEach(baseGroupItem => {
      const { template_space_id, id, name, expand, configs } = baseGroupItem
      const groupIndex = list.findIndex(item => item.id === baseGroupItem.id)
      const diffGroup: IDiffGroupData = groupIndex > -1 ? list[groupIndex] : { template_space_id, id, name, expand, configs: [] }

      configs.forEach(baseItem => {
        let currentItem: IConfigMenuItem|undefined
        currentGroupList.value.some(baseGroupItem => {
          if( baseGroupItem.template_space_id === baseGroupItem.template_space_id) {
            return baseGroupItem.configs.some((config) => {
              if (config.id === baseItem.id && config.template_revision_id === baseItem.template_revision_id) {
                currentItem = config
                return true
              }
            })
          }
          return false
        })
        if (!currentItem) {
            diffCount.value++
            diffGroup.configs.push({
                ...baseItem,
                diff_type: isBaseVersionExist.value ? 'delete' : '',
                current: '',
                base: baseItem.signature
            })
        }
      })

      if (groupIndex === -1 && diffGroup.configs.length > 0) {
        list.push(diffGroup)
      }
    })
    return list
  }

  // 设置默认选中的配置项
  // 如果props有设置选中项，取props值
  // 否则取第一个非空分组的第一个配置项
  const setDefaultSelected = () => {
    if (props.selectedConfig.id) {
      const pkg = aggregatedList.value.find(group => group.id === props.selectedConfig.pkgId)
      if (pkg) {
        pkg.expand = true
      }
      handleSelectItem(props.selectedConfig)
    } else {
      const group = aggregatedList.value.find(group => group.configs.length > 0)
      if (group) {
        handleSelectItem({ pkgId: group.id, id: group.configs[0].id, version: group.configs[0].template_revision_id })
      }
    }
  }

  const handleSearch = () => {
    if (!searchStr.value && !isOnlyShowDiff.value) {
      groupedConfigListOnShow.value = aggregatedList.value.slice()
    } else {
      const list: IDiffGroupData[] = []
      aggregatedList.value.forEach(group => {
        const configs = group.configs.filter(item => {
          console.log(item.diff_type)
          const isSearchHit = item.name.toLocaleLowerCase().includes(searchStr.value.toLocaleLowerCase())
          if (isOnlyShowDiff.value) {
            return item.diff_type !== '' && isSearchHit
          }
          return isSearchHit
        })
        if (configs.length > 0) {
          list.push({
            ...group,
            configs
          })
        }
      })
      groupedConfigListOnShow.value = list
    }
  }

  const getItemSelectedStatus = (pkgId: number, config: IConfigDiffItem) => {
    const { id, template_revision_id } = config
    return props.actived && pkgId === selected.value.pkgId && id === selected.value.id && template_revision_id === selected.value.version
  }

  // 选择对比配置项后，加载配置项详情，组装对比数据
  const handleSelectItem = async (selectedConfig: IConfigDiffSelected) => {
    const pkg = aggregatedList.value.find(item => item.id === selectedConfig.pkgId)
    if (pkg) {
      const config = pkg.configs.find(item => item.id === selectedConfig.id && item.template_revision_id === selectedConfig.version)
      if (config) {
        selected.value = selectedConfig
        const data = await getConfigDiffDetail(config)
        emits('selected', data)
      }
    }
  }

  const getConfigDiffDetail = async (config: IConfigDiffItem) => {
    let currentConfigContent: string|IFileConfigContentSummary = ''
    let baseConfigContent: string|IFileConfigContentSummary = ''
    const { id, name, file_type, update_at, byte_size, current: currentSignature, base: baseSignature } = config

    if (config.current) {
      currentConfigContent = await loadConfigContent({ id, name, file_type, update_at, byte_size, signature: currentSignature })
    }

    if (config.base) {
      baseConfigContent = await loadConfigContent({ id, name, file_type, update_at, byte_size, signature: baseSignature })
    }

    return {
      contentType: config.file_type === 'binary' ? 'file' : 'text',
      base: {
        content: baseConfigContent
      },
      current: {
        content: currentConfigContent
      }
    }
  }
  // 加载配置内容详情
  const loadConfigContent = async({ id, name, file_type, update_at, signature, byte_size }: { id: number, name: string; file_type: string; update_at: string; signature: string; byte_size: string }) => {
    if (!signature) {
      return ''
    }
    if (file_type === 'binary') {
      return {
        id,
        name,
        signature,
        update_at,
        size: byteUnitConverse(Number(byte_size))
      }
    }
    const configContent = await getConfigContent(bkBizId.value, <number>appData.value.id, signature)
    return String(configContent)
  }

</script>
<template>
  <div :class="['configs-menu', { 'search-opened':  isOpenSearch}]">
    <div class="title-area">
      <div class="title">配置项</div>
      <div class="title-extend">
        <bk-checkbox
          v-if="isBaseVersionExist"
          v-model="isOnlyShowDiff"
          class="view-diff-checkbox"
          @change="handleSearch">
          只查看差异项({{ diffCount }})
        </bk-checkbox>
        <div :class="['search-trigger', { actived: isOpenSearch }]" @click="isOpenSearch = !isOpenSearch">
          <Search />
        </div>
      </div>
    </div>
    <div v-if="isOpenSearch" class="search-wrapper">
      <SearchInput v-model="searchStr" placeholder="搜索配置项名称" @search="handleSearch" />
    </div>
    <div class="groups-wrapper">
      <div v-for="group in groupedConfigListOnShow" class="config-group-item" :key="group.id">
        <div :class="['group-header', { expand: group.expand }]" @click="group.expand = !group.expand">
          <RightShape class="arrow-icon" />
          <span v-overflow-title class="name">{{ group.name }}</span>
        </div>
        <div v-if="group.expand" class="config-list">
          <div
            v-for="config in group.configs"
            v-overflow-title
            :key="config.id"
            :class="['config-item', { actived: getItemSelectedStatus(group.id, config) }]"
            @click="handleSelectItem({ pkgId: group.id, id: config.id, version: config.template_revision_id })">
            <i v-if="config.diff_type" :class="['status-icon', config.diff_type]"></i>
            {{ config.name }}
          </div>
        </div>
      </div>
      <bk-exception
        v-if="(isOnlyShowDiff || !searchStr) && groupedConfigListOnShow.length === 0"
        class="empty-tips"
        scene="part"
        type="search-empty">
        搜索结果为空
      </bk-exception>
    </div>
  </div>
</template>
<style lang="scss" scoped>
  .configs-menu {
    background: #fafbfd;
    height: 100%;
    &.search-opened {
      .groups-wrapper {
        height: calc(100% - 80px);
      }
    }
  }
  .title-area {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 12px 8px 24px;
    .title {
      font-size: 14px;
      color: #313238;
      font-weight: 700;
    }
    .title-extend {
      display: flex;
      align-items: center;
      .view-diff-checkbox {
        padding-right: 8px;
        border-right: 1px solid #dcdee5;
        :deep(.bk-checkbox-label) {
          font-size: 12px;
        }
      }
    }
    .search-trigger {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-left: 8px;
      width: 20px;
      height: 20px;
      font-size: 12px;
      color: #63656e;
      background: #edeff1;
      border-radius: 2px;
      cursor: pointer;
      &.actived,
      &:hover {
        background: #e1ecff;
        color: #3a84ff;
      }
    }
  }
  .search-wrapper {
    padding: 0 12px 8px;
  }
  .groups-wrapper {
    height: calc(100% - 40px);
    overflow: auto;
  }
  .config-group-item {
    .group-header {
      display: flex;
      align-items: center;
      padding: 8px 12px;
      line-height: 20px;
      font-size: 12px;
      color: #313238;
      cursor: pointer;
      &.expand {
        .arrow-icon {
          transform: rotate(90deg);
          color: #3a84ff;
        }
      }
    }
    .arrow-icon {
      margin-right: 8px;
      font-size: 14px;
      color: #c4c6cc;
      transition: transform .2s ease-in-out;
    }
    .config-list {
      margin-bottom: 8px;
      .config-item {
        position: relative;
        padding: 0 12px 0 32px;
        height: 40px;
        line-height: 40px;
        font-size: 12px;
        color: #63656e;
        border-bottom: 1px solid #dcdee5;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        cursor: pointer;
        &:hover {
          background: #e1ecff;
        }
        &.actived {
          background: #e1ecff;
          color: #3a84ff;
        }
        .status-icon {
          position: absolute;
          top: 18px;
          left: 16px;
          width: 4px;
          height: 4px;
          border-radius: 50%;
          &.add {
            background: #3a84ff;
          }
          &.delete {
            background: #ea3536;
          }
          &.modify {
            background: #fe9c00;
          }
        }
      }
    }
  }
  .empty-tips {
    margin-top: 40px;
    font-size: 12px;
    color: #63656e;
  }
</style>
